<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: base.proto

namespace Base;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 *
 * IdReq 通用Id参数
 *
 * Generated from protobuf message <code>base.IdReq</code>
 */
class IdReq extends \Google\Protobuf\Internal\Message
{
    /**
     * 数据ID
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     */
    protected $id = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $id
     *           数据ID
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Base::initOnce();
        parent::__construct($data);
    }

    /**
     * 数据ID
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @return int|string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * 数据ID
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @param int|string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkInt64($var);
        $this->id = $var;

        return $this;
    }

}

