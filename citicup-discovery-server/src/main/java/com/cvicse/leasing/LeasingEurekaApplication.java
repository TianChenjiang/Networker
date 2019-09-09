package com.cvicse.leasing;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.server.EnableEurekaServer;

@SpringBootApplication
@EnableEurekaServer
public class LeasingEurekaApplication {

    public static void main(String[] args) {
        SpringApplication.run(LeasingEurekaApplication.class, args);
    }

}
