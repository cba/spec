<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: auth.proto

namespace GPBMetadata;

class Auth
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

auth.protoauth"�
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
items (2.auth.AppItem
total ("+
AppResetResp
key (	
secret (	2�
Auth-
AppList.base.ListReq.auth.AppListResp" )
	AppDetail.base.IdReq.auth.AppItem" 1
AppUpdateOrCreate.auth.AppItem.base.Empty" \'
	AppDelete.base.IdReq.base.Empty" -
AppReset.base.IdReq.auth.AppResetResp" B>
io.github.cba.authB	AuthProtoPZgithub.com/cba/spec/go/authbproto3'
        , true);

        static::$is_initialized = true;
    }
}

