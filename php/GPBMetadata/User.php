<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

namespace GPBMetadata;

class User
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Base::initOnce();
        $pool->internalAddGeneratedFile(
            '
�

user.protouser"�
AppItem

id (
identity (	
name (	
key (	
remark (	
status (

created_at (	

updated_at (	":
AppListResp
items (2.user.AppItem
total ("+
AppResetResp
key (	
secret (	2�
user-
AppList.base.ListReq.user.AppListResp" )
	AppDetail.base.IdReq.user.AppItem" 1
AppUpdateOrCreate.user.AppItem.base.Empty" \'
	AppDelete.base.IdReq.base.Empty" -
AppReset.base.IdReq.user.AppResetResp" B>
io.github.cba.userB	userProtoPZgithub.com/cba/spec/go/userbproto3'
        , true);

        static::$is_initialized = true;
    }
}

