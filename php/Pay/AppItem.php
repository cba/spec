<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pay.proto

namespace Pay;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 *
 * APP详情
 *
 * Generated from protobuf message <code>pay.AppItem</code>
 */
class AppItem extends \Google\Protobuf\Internal\Message
{
    /**
     * 数据ID,创建时传0
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     */
    protected $id = 0;
    /**
     * 权限标识
     *
     * Generated from protobuf field <code>string identity = 2;</code>
     */
    protected $identity = '';
    /**
     * APP名
     *
     * Generated from protobuf field <code>string name = 3;</code>
     */
    protected $name = '';
    /**
     * AppKey
     *
     * Generated from protobuf field <code>string key = 4;</code>
     */
    protected $key = '';
    /**
     * 备注
     *
     * Generated from protobuf field <code>string remark = 5;</code>
     */
    protected $remark = '';
    /**
     * 状态
     *
     * Generated from protobuf field <code>int64 status = 6;</code>
     */
    protected $status = 0;
    /**
     * 创建时间
     *
     * Generated from protobuf field <code>string created_at = 7;</code>
     */
    protected $created_at = '';
    /**
     * 更新时间
     *
     * Generated from protobuf field <code>string updated_at = 8;</code>
     */
    protected $updated_at = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $id
     *           数据ID,创建时传0
     *     @type string $identity
     *           权限标识
     *     @type string $name
     *           APP名
     *     @type string $key
     *           AppKey
     *     @type string $remark
     *           备注
     *     @type int|string $status
     *           状态
     *     @type string $created_at
     *           创建时间
     *     @type string $updated_at
     *           更新时间
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Pay::initOnce();
        parent::__construct($data);
    }

    /**
     * 数据ID,创建时传0
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @return int|string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * 数据ID,创建时传0
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

    /**
     * 权限标识
     *
     * Generated from protobuf field <code>string identity = 2;</code>
     * @return string
     */
    public function getIdentity()
    {
        return $this->identity;
    }

    /**
     * 权限标识
     *
     * Generated from protobuf field <code>string identity = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setIdentity($var)
    {
        GPBUtil::checkString($var, True);
        $this->identity = $var;

        return $this;
    }

    /**
     * APP名
     *
     * Generated from protobuf field <code>string name = 3;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * APP名
     *
     * Generated from protobuf field <code>string name = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * AppKey
     *
     * Generated from protobuf field <code>string key = 4;</code>
     * @return string
     */
    public function getKey()
    {
        return $this->key;
    }

    /**
     * AppKey
     *
     * Generated from protobuf field <code>string key = 4;</code>
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
     * 备注
     *
     * Generated from protobuf field <code>string remark = 5;</code>
     * @return string
     */
    public function getRemark()
    {
        return $this->remark;
    }

    /**
     * 备注
     *
     * Generated from protobuf field <code>string remark = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setRemark($var)
    {
        GPBUtil::checkString($var, True);
        $this->remark = $var;

        return $this;
    }

    /**
     * 状态
     *
     * Generated from protobuf field <code>int64 status = 6;</code>
     * @return int|string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * 状态
     *
     * Generated from protobuf field <code>int64 status = 6;</code>
     * @param int|string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkInt64($var);
        $this->status = $var;

        return $this;
    }

    /**
     * 创建时间
     *
     * Generated from protobuf field <code>string created_at = 7;</code>
     * @return string
     */
    public function getCreatedAt()
    {
        return $this->created_at;
    }

    /**
     * 创建时间
     *
     * Generated from protobuf field <code>string created_at = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setCreatedAt($var)
    {
        GPBUtil::checkString($var, True);
        $this->created_at = $var;

        return $this;
    }

    /**
     * 更新时间
     *
     * Generated from protobuf field <code>string updated_at = 8;</code>
     * @return string
     */
    public function getUpdatedAt()
    {
        return $this->updated_at;
    }

    /**
     * 更新时间
     *
     * Generated from protobuf field <code>string updated_at = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setUpdatedAt($var)
    {
        GPBUtil::checkString($var, True);
        $this->updated_at = $var;

        return $this;
    }

}

