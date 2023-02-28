<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: order.proto

namespace Order;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 *
 * APP重置返回
 *
 * Generated from protobuf message <code>order.AppResetResp</code>
 */
class AppResetResp extends \Google\Protobuf\Internal\Message
{
    /**
     * AppKey
     *
     * Generated from protobuf field <code>string key = 1;</code>
     */
    protected $key = '';
    /**
     * AppSecret
     *
     * Generated from protobuf field <code>string secret = 2;</code>
     */
    protected $secret = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $key
     *           AppKey
     *     @type string $secret
     *           AppSecret
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Order::initOnce();
        parent::__construct($data);
    }

    /**
     * AppKey
     *
     * Generated from protobuf field <code>string key = 1;</code>
     * @return string
     */
    public function getKey()
    {
        return $this->key;
    }

    /**
     * AppKey
     *
     * Generated from protobuf field <code>string key = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->key = $var;

        return $this;
    }

    /**
     * AppSecret
     *
     * Generated from protobuf field <code>string secret = 2;</code>
     * @return string
     */
    public function getSecret()
    {
        return $this->secret;
    }

    /**
     * AppSecret
     *
     * Generated from protobuf field <code>string secret = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setSecret($var)
    {
        GPBUtil::checkString($var, True);
        $this->secret = $var;

        return $this;
    }

}
