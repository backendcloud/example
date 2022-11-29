package com.example.demo.controller;

import com.example.demo.service.AsyncHasReturn;
import com.example.demo.service.AsyncNoReturn;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.concurrent.ExecutionException;
import java.util.concurrent.Future;

@RestController
@RequestMapping()
// @RequiresAuthentication
public class DemoController {

    private AsyncNoReturn asyncNoReturn;
    @Autowired
    public void setService(AsyncNoReturn asyncNoReturn) {
        this.asyncNoReturn = asyncNoReturn;
    }

    private AsyncHasReturn asyncHasReturn;
    @Autowired
    public void setService(AsyncHasReturn asyncHasReturn) {
        this.asyncHasReturn = asyncHasReturn;
    }

    @GetMapping("/async-no-return")
    public ResponseEntity<String> getAsyncNoReturn() {
        long s = System.currentTimeMillis();
        asyncNoReturn.execAsync1();
        asyncNoReturn.execAsync2();
        asyncNoReturn.execAsync3();
        System.out.println("我执行结束了");
        long costTime = System.currentTimeMillis() - s;
        System.out.println("service async-no-return cost time: " + costTime);
        return ResponseEntity.status(200).body("good");
    }

    @GetMapping("/async-has-return")
    public ResponseEntity<String> getAsyncHasReturn() throws ExecutionException, InterruptedException {
        long s = System.currentTimeMillis();
        Future<String> stringFuture1 = asyncHasReturn.execAsync1();
        Future<String> stringFuture2 = asyncHasReturn.execAsync2();
        Future<String> stringFuture3 = asyncHasReturn.execAsync3();
        long costTime = System.currentTimeMillis() - s;
        System.out.println("执行到代码中间了，cost time: " + costTime);
        System.out.println("我执行结束了" + stringFuture1.get() + stringFuture2.get() + stringFuture3.get());
        costTime = System.currentTimeMillis() - s;
        System.out.println("service async-has-return cost time: " + costTime);
        return ResponseEntity.status(200).body("good");
    }

}