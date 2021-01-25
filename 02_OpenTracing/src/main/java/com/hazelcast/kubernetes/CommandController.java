package com.hazelcast.kubernetes;

 /* ~~ OpenTracing ~~
import io.opentracing.contrib.hazelcast.TracingHazelcastInstance;
import io.opentracing.Tracer;
import io.opentracing.util.GlobalTracer;
 */
import com.hazelcast.client.HazelcastClient;
import com.hazelcast.core.HazelcastInstance;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.concurrent.ConcurrentMap;

import java.net.*; 
import java.io.*; 
import java.util.*; 
import java.net.InetAddress; 

@RestController
public class CommandController {

    @Autowired
    HazelcastInstance hazelcastInstance;

     /* ~~ OpenTracing ~~
    public Tracer initTracer() {
        return GlobalTracer.get();
    }
     */
    
    public HazelcastInstance getTracingHazelcast() {
         /* ~~ OpenTracing ~~
        if (hazelcastInstance instanceof TracingHazelcastInstance) {
            return hazelcastInstance;
        }
        hazelcastInstance = new TracingHazelcastInstance(
                HazelcastClient.newHazelcastClient(),
                false);
         */
        return hazelcastInstance;
    }

    @RequestMapping("/put")
    public CommandResponse put(@RequestParam(value = "key") String key, @RequestParam(value = "value") String value) {
        ConcurrentMap<String, String> map = getTracingHazelcast().getMap("map");
        String oldValue = map.put(key, value);
        return new CommandResponse(oldValue);
    }

    @RequestMapping("/get")
    public CommandResponse get(@RequestParam(value = "key") String key) {
        ConcurrentMap<String, String> map = getTracingHazelcast().getMap("map");
        String value = map.get(key);
        return new CommandResponse(value);
    }

    @RequestMapping("/remove")
    public CommandResponse remove(@RequestParam(value = "key") String key) {
        ConcurrentMap<String, String> map = getTracingHazelcast().getMap("map");
        String value = map.remove(key);
        return new CommandResponse(value);
    }

    @RequestMapping("/size")
    public CommandResponse size() {
        ConcurrentMap<String, String> map = getTracingHazelcast().getMap("map");
        int size = map.size();
        return new CommandResponse(Integer.toString(size));
    }

    @RequestMapping("/pop1")
    public CommandResponse pop1() {
        ConcurrentMap<String, String> map = getTracingHazelcast().getMap("map");
        String addr = "";
        try {
            URL url_name = new URL("http://ifconfig.me/ip"); 
            BufferedReader sc = new BufferedReader(new InputStreamReader(url_name.openStream())); 
            addr = sc.readLine().trim(); 
        }
        catch (Exception e) 
        { 
            System.out.println("Error" + e); 
        } 

        for (int i = 0; i < 100; i++) {
            String s = addr + "-" + Integer.toString(i);
            map.put(s, s);
        }
        return new CommandResponse("100 entry inserted to the map with key: " + addr + "-* , starting with 1 ");
    }
 
    @RequestMapping("/pop2")
    public CommandResponse pop2() {
        ConcurrentMap<String, String> map = getTracingHazelcast().getMap("map");
        String addr = "";
        try {
            URL url_name = new URL("http://ifconfig.me/ip"); 
            BufferedReader sc = new BufferedReader(new InputStreamReader(url_name.openStream())); 
            addr = sc.readLine().trim(); 
        }
        catch (Exception e) 
        { 
            System.out.println("Error" + e); 
        } 

        for (int i = 200; i < 300; i++) {
            String s = addr + "-" + Integer.toString(i);
            map.put(s, s);
        }
        return new CommandResponse("100 entry inserted to the map with key: " + addr + "-2** , starting with 200 ");
    }

    @RequestMapping("/clear")
    public CommandResponse clear() {
        ConcurrentMap<String, String> map = getTracingHazelcast().getMap("map");
        map.clear();
        return new CommandResponse("Map cleared");
    }
}
