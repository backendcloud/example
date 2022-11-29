package com.example.demo.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
public class QuickStartController {
    @RequestMapping("/test")
    @ResponseBody
    public String test(){
        return "springboot 3.0 访问测试";
    }

    @RequestMapping("/hello")
    @ResponseBody
    public String home(){
        return "Hello World from springboot 3.0!";
    }
}
