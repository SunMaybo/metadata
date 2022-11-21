package com.jewel.meta.asset_platform.proto.greeter_service;/**
 * @Description:
 * @Author: fico
 * @Date: 2022/11/21
 **/

import io.grpc.Channel;
import io.grpc.ManagedChannelBuilder;

import java.util.Arrays;
import java.util.concurrent.TimeUnit;

/**
 * @Description:
 * @Author: fico
 * @Date: 2022/11/21
 **/
public class main {

    public static void main(String[] args) throws InterruptedException {
        Channel ch = ManagedChannelBuilder.forTarget("192.168.0.103:8088").usePlaintext().keepAliveTime(10, TimeUnit.SECONDS).keepAliveTimeout(1, TimeUnit.SECONDS).usePlaintext().enableRetry().maxRetryAttempts(3).build();
        GreeterServiceGrpc.newBlockingStub(ch).withWaitForReady().withDeadlineAfter(5, TimeUnit.SECONDS).sayHelloWord(GreeterServiceProtocol.HelloRequest.newBuilder().setName("xxxxx").build());
        Thread.sleep(10000);
        long start=System.currentTimeMillis();
        try {
            GreeterServiceGrpc.newBlockingStub(ch).withWaitForReady().withDeadlineAfter(5, TimeUnit.SECONDS).sayHelloWord(GreeterServiceProtocol.HelloRequest.newBuilder().setName("xxxxx").build());
        }catch (Exception e){
            throw new RuntimeException(e);
        }finally {
            System.out.println("args = " + (System.currentTimeMillis()-start));
        }
    }
}
