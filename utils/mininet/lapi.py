from mn_wifi.net import Mininet_wifi
from mn_wifi.node import AP, Station, Car, OVSKernelAP, physicalAP
from mininet.log import info
from os import environ
import threading
import time
import json
import datetime

try:
  from urllib.request import build_opener, HTTPHandler, Request
except ImportError:
  from urllib2 import build_opener, HTTPHandler, Request

LAPI_URL = environ.get('LAPI_URL', 'http://192.168.124.1:8080/mininet')
USER_REGISTER_URL = LAPI_URL + '/user'
USER_LOCATION_URL = LAPI_URL + '/location'
AP_REGISTER_URL = LAPI_URL + '/ap'
AP_LOCATION_URL = LAPI_URL + '/aplocation'

# Global variable to store the state of export_users
export_users_done = False
export_aps_done = False

class LapiUser:
    """ Represents a User in the Location API"""
    def __init__(self, address, access_point):
        self.address = address
        self.access_point = access_point
        
    def to_dict(self):
        return {
            "address": self.address,
            "access_point": self.access_point
        }
    
    def to_json(self):
        return json.dumps(self.to_dict)

class LapiUserLocation:
    """
    Represents a user with location information for LAPI integration
    """
    def __init__(self, userid, coordx, coordy, coordz, apname=None):
        """
        Initialize a LapiUser object
        
        Args:
            userid (str): User identifier
            coordx (float): X coordinate position
            coordy (float): Y coordinate position
            coordz (float): Z coordinate position
            apname (str, optional): Name of the AP the user is connected to
            timestamp (str, optional): Timestamp of the location data
        """
        self.userid = userid
        self.coordx = float(coordx)
        self.coordy = float(coordy)
        self.coordz = float(coordz)
        self.apname = apname
        self.timestamp = int(time.time())
    
    def to_dict(self):
        """
        Convert the LapiUser object to a dictionary
        
        Returns:
            dict: Dictionary representation of the LapiUser
        """
        return {
            "userid": self.userid,
            "coordx": self.coordx,
            "coordy": self.coordy,
            "coordz": self.coordz,
            "apname": self.apname,
            "timestamp": self.timestamp
        }
    
    def to_json(self):
        """
        Convert the LapiUser object to a JSON string
        
        Returns:
            str: JSON string representation of the LapiUser
        """
        return json.dumps(self.to_dict())

class LapiApLocation:
    """
    Represents an Access Point (AP) with location information for LAPI integration
    """
    def __init__(self, apid, coordx, coordy, coordz):
        """
        Initialize a LapiAP object
        
        Args:
            apid (str): Identifier of the AP
            coordx (float): X coordinate position
            coordy (float): Y coordinate position
            coordz (float): Z coordinate position
            timestamp (str, optional): Timestamp of the location data
        """
        self.apid = apid
        self.coordx = float(coordx)
        self.coordy = float(coordy)
        self.coordz = float(coordz)
        self.timestamp = int(time.time())
    
    def to_dict(self):
        """
        Convert the LapiAP object to a dictionary
        
        Returns:
            dict: Dictionary representation of the LapiAP
        """
        return {
            "apid": self.apid,
            "coordx": self.coordx,
            "coordy": self.coordy,
            "coordz": self.coordz,
            "timestamp": self.timestamp
        }
    
    def to_json(self):
        """
        Convert the LapiAP object to a JSON string
        
        Returns:
            str: JSON string representation of the LapiAP
        """
        return json.dumps(self.to_dict())

def wifi_wrapper(fn):

    def lapi_post(url, data):
        """
        Send a POST request to the LAPI server
        
        Args:
            url (str): URL of the LAPI server
            data (dict): Data to be sent in the request
        """
        try:
            opener = build_opener(HTTPHandler)
            request = Request(url, data=data.encode('utf-8'), headers={'Content-Type': 'application/json'})
            info(f"Sending LAPI request to {url} with data: {data}\n")
            response = opener.open(request)
            return response.read()
        except Exception as e:
            print(f"Error sending LAPI request: {e}\n")
            return None


    def get_ap_name(node):
        """ Gets the Access Point name which a node is connected to."""
        if node.params['wlan'][0]:
            wintf  = node.getNameToWintf(node.params['wlan'][0])
            wintf_name = wintf.associatedTo if wintf.associatedTo else None
            ap_name = str(wintf_name).split('-')[0] if wintf_name else None
        else: 
            ap_name = None
        return ap_name    
    

    def export_users(net):
        """ Export all users (stations) in the Mininet network at the 
            beggining of the experiment. This function should be called before
            monitor_locations() and run once. A global variable is used to
            store the state of this function.
        """
        global export_users_done
        if not export_users_done:
            export_users_done = True
            info("*** Exporting users\n")
            nodes = net.get_mn_wifi_nodes()
            users = [n for n in nodes if isinstance(n, (Station, Car))]
            # send the list of users to the server via API
            for u in users:
                address = u.params['ip'].split('/')[0]
                access_point = get_ap_name(u)
                lapi_post(USER_REGISTER_URL, LapiUser(address, access_point).to_json())

    def export_aps(net):
        """ Export all Access Points (APs) in the Mininet network at the 
            beggining of the experiment. This function should be called before
            monitor_locations() and run once. A global variable is used to
            store the state of this function.
        """
        global export_aps_done
        if not export_aps_done:
            export_aps_done = True
            info("*** Exporting Access Points\n")
            nodes = net.get_mn_wifi_nodes()
            aps = [n for n in nodes if isinstance(n, (AP))]
            for ap in aps:
                (x, y, z) = ap.getxyz()
                lapi_post(AP_REGISTER_URL, LapiApLocation(ap.name, x, y, z).to_json())


    def monitor_locations(net, interval=5):
        """Periodically sends location information of stations and APs"""
        
        def location_monitor():
            while True:
                try:
                    print("\n*** Current Network Locations:\n")
                    # save location to a file
                    nodes = net.get_mn_wifi_nodes()
                    for n in nodes:
                        if isinstance(n, (Station, Car)):
                            (x, y, z) = n.getxyz()
                            if n.params['wlan'][0]: 
                                ap_name = get_ap_name(n)
                            else:
                                ap_name = None
                            # info(f"Station {n.name} (AP: {ap_name}) location: {x}, {y}, {z}\n")
                            # Send user data to LAPI
                            node_ip = n.params['ip'].split('/')[0]
                            # user = LapiUser(n.name, x, y, z, ap_name)
                            user = LapiUserLocation(node_ip, x, y, z, ap_name)
                            data = user.to_json()
                            lapi_post(USER_LOCATION_URL, data)
                        if isinstance(n, (AP)):
                            (x, y, z) = n.getxyz()
                            ap = LapiApLocation(n.name, x, y, z)
                            data = ap.to_json()
                            lapi_post(AP_LOCATION_URL, data)
                    time.sleep(interval)
                except Exception as e:
                    print(f"Error in location monitoring: {e}\n")
                    break
        
        # Start monitoring in a separate thread to avoid blocking
        monitor_thread = threading.Thread(target=location_monitor)
        monitor_thread.daemon = True  # Thread will exit when main program exits
        monitor_thread.start()
    
    def result(*args, **kwargs):
        res = fn(*args, **kwargs)
        net = args[0]

        # Export all the users (stations) in the Mininet network
        export_users(net)
        # Export all the Access Points (APs) in the Mininet network
        export_aps(net)
        
        # Add users from Mininet_wifi
        send_users(net)
        # Start location monitoring with default 5-second interval
        # (can be adjusted by setting LOCATION_INTERVAL env var)
        interval = int(environ.get('LOCATION_INTERVAL', '1'))
        monitor_locations(net, interval)
        
        return res
    
    return result

# Import Mininet_wifi and hook the start method
try:
    setattr(Mininet_wifi, 'build', wifi_wrapper(Mininet_wifi.__dict__['build']))
    print("*** Successfully hooked Mininet_wifi for location monitoring\n")
except ImportError:
    print("*** Error: location monitoring not enabled\n")
