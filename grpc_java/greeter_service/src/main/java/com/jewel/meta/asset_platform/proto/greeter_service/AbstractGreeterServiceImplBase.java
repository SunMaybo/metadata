package com.jewel.meta.asset_platform.proto.greeter_service;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.util.JsonFormat;
import io.grpc.Context;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public abstract class AbstractGreeterServiceImplBase extends GreeterServiceGrpc.GreeterServiceImplBase {
   private  final Logger LOGGER= LoggerFactory.getLogger(this.getClass());
   
	
    @Override
    public void  sayHelloWord(com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloRequest request, io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloReply> responseObserver) {
 		try {        
 			if (Context.current().isCancelled()) {
                responseObserver.onError(Status.CANCELLED.withDescription("Cancelled by client").asRuntimeException());
                return;
            }
			responseObserver.onNext(sayHelloWord(request));
        	responseObserver.onCompleted();
        } catch (StatusRuntimeException e) {
			try {
                LOGGER.warn("Caller Request sayHelloWord======>" + JsonFormat.printer().print(request));
            } catch (InvalidProtocolBufferException ex) {
               	LOGGER.error("grpc caller err:",e);
            }
			LOGGER.warn("grpc caller err:",e);
            responseObserver.onError(e);
        }
    }
    /**
     */
    protected abstract com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloReply sayHelloWord(com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloRequest request);
   
   
	
    @Override
    public io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloRequest> sayStream(io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloReply> responseObserver) {
        return sayStreamWithDuplex(responseObserver);
    }
    /**
     */
    protected abstract io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloRequest> sayStreamWithDuplex(io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloReply> responseObserver);

   
   
	
    @Override
    public io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloRequest> sayStream1(io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloReply> responseObserver) {
        return sayStream1WithDuplex(responseObserver);
    }
    /**
     */
    protected abstract io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloRequest> sayStream1WithDuplex(io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloReply> responseObserver);

   
   
	
    @Override
    public void sayStream2(com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloRequest request, io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloReply> responseObserver) {
        sayStream2WithReturnSimplex(request, responseObserver);
    }
	/**
     */
    public abstract void sayStream2WithReturnSimplex(com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloRequest request, io.grpc.stub.StreamObserver<com.jewel.meta.asset_platform.proto.greeter_service.GreeterServiceProtocol.HelloReply> responseObserver);
   
   
}
