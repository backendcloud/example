package com.example.demo.service.impl;

import com.example.demo.service.AsyncNoReturn;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;

@Service
public class AsyncNoReturnImpl implements AsyncNoReturn {

    @Async
    @Override
    public void execAsync1(){
        try{
            Thread.sleep(2000);
            System.out.println("甲睡了2000ms");
        }catch (Exception e){
            e.printStackTrace();
        }
    }
    @Async
    @Override
    public void execAsync2(){
        try{
            Thread.sleep(4000);
            System.out.println("乙睡了4000ms");
        }catch (Exception e){
            e.printStackTrace();
        }
    }
    @Async
    @Override
    public void execAsync3(){
        try{
            Thread.sleep(3000);
            System.out.println("丙睡了3000ms");
        }catch (Exception e){
            e.printStackTrace();
        }
    }
}
