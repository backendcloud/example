package com.example.demo.service.impl;

import com.example.demo.service.AsyncHasReturn;
import org.springframework.scheduling.annotation.Async;
import org.springframework.scheduling.annotation.AsyncResult;
import org.springframework.stereotype.Service;

import java.util.concurrent.Future;

@Service
public class AsyncHasReturnImpl implements AsyncHasReturn {

    @Async
    public Future<String> execAsync1(){
        try{
            Thread.sleep(2000);
            System.out.println("甲睡了2000ms");
            return new AsyncResult<String>("甲执行成功了");
        }catch (Exception e){
            e.printStackTrace();
        }
        return null;
    }
    @Async
    public Future<String> execAsync2(){
        try{
            Thread.sleep(4000);
            System.out.println("乙睡了4000ms");
            return new AsyncResult<String>("乙执行成功了");
        }catch (Exception e){
            e.printStackTrace();
        }
        return null;
    }
    @Async
    public Future<String> execAsync3(){
        try{
            Thread.sleep(3000);
            System.out.println("丙睡了3000ms");
            return new  AsyncResult<String>("丙执行结束了");
        }catch (Exception e){
            e.printStackTrace();
        }
        return null;
    }

}
