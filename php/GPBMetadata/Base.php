<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: base.proto

namespace GPBMetadata;

class Base
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�

base.protobase"
Empty"4
ListReq
cond (	
page (
limit ("
IdReq

id ("U
Map!
data (2.base.Map.DataEntry+
	DataEntry
key (	
value (	:8B>
io.github.cba.baseB	BaseProtoPZgithub.com/cba/spec/go/basebproto3'
        , true);

        static::$is_initialized = true;
    }
}

