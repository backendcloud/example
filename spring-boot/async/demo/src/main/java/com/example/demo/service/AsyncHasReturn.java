package com.example.demo.service;


import java.util.concurrent.Future;

public interface AsyncHasReturn {
    Future<String> execAsync1();
    Future<String> execAsync2();
    Future<String> execAsync3();
}
