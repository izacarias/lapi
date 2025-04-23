from mn_wifi.net import Mininet_wifi
from mn_wifi.node import AP, Station, Car, OVSKernelAP, physicalAP
from mininet.log import info
from os import environ
import threading
import time

try:
  from urllib.request import build_opener, HTTPHandler, Request
except ImportError:
  from urllib2 import build_opener, HTTPHandler, Request

def wifi_wrapper(fn):
    
    def monitor_locations(net, interval=5):
        """Periodically prints location information of stations and APs"""
        
        def location_monitor():
            while True:
                try:
                    print("\n*** Current Network Locations:\n")
                    # save location to a file
                    nodes = net.get_mn_wifi_nodes()
                    for n in nodes:
                        if isinstance(n, (Station, Car)):
                            (x, y, z) = n.getxyz()
                            info(f"Station {n.name} location: {x}, {y}, {z}\n")
                        if isinstance(n, (AP)):
                            (x, y, z) = n.getxyz()
                            info(f"AP {n.name} location: {x}, {y}, {z}\n")
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
        
        # Start location monitoring with default 5-second interval
        # (can be adjusted by setting LOCATION_INTERVAL env var)
        interval = int(environ.get('LOCATION_INTERVAL', '5'))
        monitor_locations(net, interval)
        
        return res
    
    return result

# Import Mininet_wifi and hook the start method
try:
    setattr(Mininet_wifi, 'build', wifi_wrapper(Mininet_wifi.__dict__['build']))
    print("*** Successfully hooked Mininet_wifi for location monitoring\n")
except ImportError:
    print("*** Error: location monitoring not enabled\n")
