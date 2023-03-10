<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pay.proto

namespace Pay;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 *
 * 生成支付单返回值
 *
 * Generated from protobuf message <code>pay.GenResp</code>
 */
class GenResp extends \Google\Protobuf\Internal\Message
{
    /**
     * 支付单号
     *
     * Generated from protobuf field <code>string payment_no = 1;</code>
     */
    protected $payment_no = '';
    /**
     * 订单号
     *
     * Generated from protobuf field <code>string order_no = 2;</code>
     */
    protected $order_no = '';
    /**
     * 支持的支付方式列表
     *
     * Generated from protobuf field <code>repeated .pay.payType pay_type = 3;</code>
     */
    private $pay_type;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $payment_no
     *           支付单号
     *     @type string $order_no
     *           订单号
     *     @type array<\Pay\payType>|\Google\Protobuf\Internal\RepeatedField $pay_type
     *           支持的支付方式列表
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Pay::initOnce();
        parent::__construct($data);
    }

    /**
     * 支付单号
     *
     * Generated from protobuf field <code>string payment_no = 1;</code>
     * @return string
     */
    public function getPaymentNo()
    {
        return $this->payment_no;
    }

    /**
     * 支付单号
     *
     * Generated from protobuf field <code>string payment_no = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setPaymentNo($var)
    {
        GPBUtil::checkString($var, True);
        $this->payment_no = $var;

        return $this;
    }

    /**
     * 订单号
     *
     * Generated from protobuf field <code>string order_no = 2;</code>
     * @return string
     */
    public function getOrderNo()
    {
        return $this->order_no;
    }

    /**
     * 订单号
     *
     * Generated from protobuf field <code>string order_no = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setOrderNo($var)
    {
        GPBUtil::checkString($var, True);
        $this->order_no = $var;

        return $this;
    }

    /**
     * 支持的支付方式列表
     *
     * Generated from protobuf field <code>repeated .pay.payType pay_type = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getPayType()
    {
        return $this->pay_type;
    }

    /**
     * 支持的支付方式列表
     *
     * Generated from protobuf field <code>repeated .pay.payType pay_type = 3;</code>
     * @param array<\Pay\payType>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setPayType($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Pay\payType::class);
        $this->pay_type = $arr;

        return $this;
    }

}

