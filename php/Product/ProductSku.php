<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: product.proto

namespace Product;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>product.ProductSku</code>
 */
class ProductSku extends \Google\Protobuf\Internal\Message
{
    /**
     * 数据ID,创建时传0
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     */
    protected $id = 0;
    /**
     * attrs md5后的字符串
     *
     * Generated from protobuf field <code>string spec_no = 2;</code>
     */
    protected $spec_no = '';
    /**
     * 销售属性值{attr_id}:{attr_value},{attr_id}:{attr_value} 多个销售属性值逗号分隔
     *
     * Generated from protobuf field <code>repeated .product.ProductAttr attrs = 3;</code>
     */
    private $attrs;
    /**
     * 商品介绍主图 多个图片逗号分隔
     *
     * Generated from protobuf field <code>string cover_url = 4;</code>
     */
    protected $cover_url = '';
    /**
     * 库存
     *
     * Generated from protobuf field <code>int64 in_stock = 5;</code>
     */
    protected $in_stock = 0;
    /**
     * 销量
     *
     * Generated from protobuf field <code>int64 sales = 6;</code>
     */
    protected $sales = 0;
    /**
     * 自定义标签
     *
     * Generated from protobuf field <code>string custom_label = 7;</code>
     */
    protected $custom_label = '';
    /**
     * 售价，整数方式保存
     *
     * Generated from protobuf field <code>int64 price_fee = 8;</code>
     */
    protected $price_fee = 0;
    /**
     * 市场价，整数方式保存
     *
     * Generated from protobuf field <code>int64 consign_price = 9;</code>
     */
    protected $consign_price = 0;
    /**
     * 成本价，整数方式保存
     *
     * Generated from protobuf field <code>int64 market_price_fee = 10;</code>
     */
    protected $market_price_fee = 0;
    /**
     * 状态 1:enable, 0:disable
     *
     * Generated from protobuf field <code>int64 status = 11;</code>
     */
    protected $status = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $id
     *           数据ID,创建时传0
     *     @type string $spec_no
     *           attrs md5后的字符串
     *     @type array<\Product\ProductAttr>|\Google\Protobuf\Internal\RepeatedField $attrs
     *           销售属性值{attr_id}:{attr_value},{attr_id}:{attr_value} 多个销售属性值逗号分隔
     *     @type string $cover_url
     *           商品介绍主图 多个图片逗号分隔
     *     @type int|string $in_stock
     *           库存
     *     @type int|string $sales
     *           销量
     *     @type string $custom_label
     *           自定义标签
     *     @type int|string $price_fee
     *           售价，整数方式保存
     *     @type int|string $consign_price
     *           市场价，整数方式保存
     *     @type int|string $market_price_fee
     *           成本价，整数方式保存
     *     @type int|string $status
     *           状态 1:enable, 0:disable
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Product::initOnce();
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
     * attrs md5后的字符串
     *
     * Generated from protobuf field <code>string spec_no = 2;</code>
     * @return string
     */
    public function getSpecNo()
    {
        return $this->spec_no;
    }

    /**
     * attrs md5后的字符串
     *
     * Generated from protobuf field <code>string spec_no = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setSpecNo($var)
    {
        GPBUtil::checkString($var, True);
        $this->spec_no = $var;

        return $this;
    }

    /**
     * 销售属性值{attr_id}:{attr_value},{attr_id}:{attr_value} 多个销售属性值逗号分隔
     *
     * Generated from protobuf field <code>repeated .product.ProductAttr attrs = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAttrs()
    {
        return $this->attrs;
    }

    /**
     * 销售属性值{attr_id}:{attr_value},{attr_id}:{attr_value} 多个销售属性值逗号分隔
     *
     * Generated from protobuf field <code>repeated .product.ProductAttr attrs = 3;</code>
     * @param array<\Product\ProductAttr>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAttrs($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Product\ProductAttr::class);
        $this->attrs = $arr;

        return $this;
    }

    /**
     * 商品介绍主图 多个图片逗号分隔
     *
     * Generated from protobuf field <code>string cover_url = 4;</code>
     * @return string
     */
    public function getCoverUrl()
    {
        return $this->cover_url;
    }

    /**
     * 商品介绍主图 多个图片逗号分隔
     *
     * Generated from protobuf field <code>string cover_url = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setCoverUrl($var)
    {
        GPBUtil::checkString($var, True);
        $this->cover_url = $var;

        return $this;
    }

    /**
     * 库存
     *
     * Generated from protobuf field <code>int64 in_stock = 5;</code>
     * @return int|string
     */
    public function getInStock()
    {
        return $this->in_stock;
    }

    /**
     * 库存
     *
     * Generated from protobuf field <code>int64 in_stock = 5;</code>
     * @param int|string $var
     * @return $this
     */
    public function setInStock($var)
    {
        GPBUtil::checkInt64($var);
        $this->in_stock = $var;

        return $this;
    }

    /**
     * 销量
     *
     * Generated from protobuf field <code>int64 sales = 6;</code>
     * @return int|string
     */
    public function getSales()
    {
        return $this->sales;
    }

    /**
     * 销量
     *
     * Generated from protobuf field <code>int64 sales = 6;</code>
     * @param int|string $var
     * @return $this
     */
    public function setSales($var)
    {
        GPBUtil::checkInt64($var);
        $this->sales = $var;

        return $this;
    }

    /**
     * 自定义标签
     *
     * Generated from protobuf field <code>string custom_label = 7;</code>
     * @return string
     */
    public function getCustomLabel()
    {
        return $this->custom_label;
    }

    /**
     * 自定义标签
     *
     * Generated from protobuf field <code>string custom_label = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomLabel($var)
    {
        GPBUtil::checkString($var, True);
        $this->custom_label = $var;

        return $this;
    }

    /**
     * 售价，整数方式保存
     *
     * Generated from protobuf field <code>int64 price_fee = 8;</code>
     * @return int|string
     */
    public function getPriceFee()
    {
        return $this->price_fee;
    }

    /**
     * 售价，整数方式保存
     *
     * Generated from protobuf field <code>int64 price_fee = 8;</code>
     * @param int|string $var
     * @return $this
     */
    public function setPriceFee($var)
    {
        GPBUtil::checkInt64($var);
        $this->price_fee = $var;

        return $this;
    }

    /**
     * 市场价，整数方式保存
     *
     * Generated from protobuf field <code>int64 consign_price = 9;</code>
     * @return int|string
     */
    public function getConsignPrice()
    {
        return $this->consign_price;
    }

    /**
     * 市场价，整数方式保存
     *
     * Generated from protobuf field <code>int64 consign_price = 9;</code>
     * @param int|string $var
     * @return $this
     */
    public function setConsignPrice($var)
    {
        GPBUtil::checkInt64($var);
        $this->consign_price = $var;

        return $this;
    }

    /**
     * 成本价，整数方式保存
     *
     * Generated from protobuf field <code>int64 market_price_fee = 10;</code>
     * @return int|string
     */
    public function getMarketPriceFee()
    {
        return $this->market_price_fee;
    }

    /**
     * 成本价，整数方式保存
     *
     * Generated from protobuf field <code>int64 market_price_fee = 10;</code>
     * @param int|string $var
     * @return $this
     */
    public function setMarketPriceFee($var)
    {
        GPBUtil::checkInt64($var);
        $this->market_price_fee = $var;

        return $this;
    }

    /**
     * 状态 1:enable, 0:disable
     *
     * Generated from protobuf field <code>int64 status = 11;</code>
     * @return int|string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * 状态 1:enable, 0:disable
     *
     * Generated from protobuf field <code>int64 status = 11;</code>
     * @param int|string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkInt64($var);
        $this->status = $var;

        return $this;
    }

}

