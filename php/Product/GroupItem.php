<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: product.proto

namespace Product;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 *
 * Group 详情
 *
 * Generated from protobuf message <code>product.GroupItem</code>
 */
class GroupItem extends \Google\Protobuf\Internal\Message
{
    /**
     * 数据ID,创建时传0
     *
     * Generated from protobuf field <code>int64 id = 1;</code>
     */
    protected $id = 0;
    /**
     * 分组类型 base基础属性, sale销售属性, categ类别属性
     *
     * Generated from protobuf field <code>string group_type = 2;</code>
     */
    protected $group_type = '';
    /**
     * 分组名称
     *
     * Generated from protobuf field <code>string name = 3;</code>
     */
    protected $name = '';
    /**
     * 分组描述
     *
     * Generated from protobuf field <code>string desc = 4;</code>
     */
    protected $desc = '';
    /**
     * 自定义标签
     *
     * Generated from protobuf field <code>string custom_label = 5;</code>
     */
    protected $custom_label = '';
    /**
     * 排序
     *
     * Generated from protobuf field <code>int64 sort = 6;</code>
     */
    protected $sort = 0;
    /**
     * 状态
     *
     * Generated from protobuf field <code>int64 status = 7;</code>
     */
    protected $status = 0;
    /**
     * 属性列表
     *
     * Generated from protobuf field <code>repeated .product.AttrItem items = 8;</code>
     */
    private $items;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $id
     *           数据ID,创建时传0
     *     @type string $group_type
     *           分组类型 base基础属性, sale销售属性, categ类别属性
     *     @type string $name
     *           分组名称
     *     @type string $desc
     *           分组描述
     *     @type string $custom_label
     *           自定义标签
     *     @type int|string $sort
     *           排序
     *     @type int|string $status
     *           状态
     *     @type array<\Product\AttrItem>|\Google\Protobuf\Internal\RepeatedField $items
     *           属性列表
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
     * 分组类型 base基础属性, sale销售属性, categ类别属性
     *
     * Generated from protobuf field <code>string group_type = 2;</code>
     * @return string
     */
    public function getGroupType()
    {
        return $this->group_type;
    }

    /**
     * 分组类型 base基础属性, sale销售属性, categ类别属性
     *
     * Generated from protobuf field <code>string group_type = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setGroupType($var)
    {
        GPBUtil::checkString($var, True);
        $this->group_type = $var;

        return $this;
    }

    /**
     * 分组名称
     *
     * Generated from protobuf field <code>string name = 3;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * 分组名称
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
     * 分组描述
     *
     * Generated from protobuf field <code>string desc = 4;</code>
     * @return string
     */
    public function getDesc()
    {
        return $this->desc;
    }

    /**
     * 分组描述
     *
     * Generated from protobuf field <code>string desc = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setDesc($var)
    {
        GPBUtil::checkString($var, True);
        $this->desc = $var;

        return $this;
    }

    /**
     * 自定义标签
     *
     * Generated from protobuf field <code>string custom_label = 5;</code>
     * @return string
     */
    public function getCustomLabel()
    {
        return $this->custom_label;
    }

    /**
     * 自定义标签
     *
     * Generated from protobuf field <code>string custom_label = 5;</code>
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
     * 排序
     *
     * Generated from protobuf field <code>int64 sort = 6;</code>
     * @return int|string
     */
    public function getSort()
    {
        return $this->sort;
    }

    /**
     * 排序
     *
     * Generated from protobuf field <code>int64 sort = 6;</code>
     * @param int|string $var
     * @return $this
     */
    public function setSort($var)
    {
        GPBUtil::checkInt64($var);
        $this->sort = $var;

        return $this;
    }

    /**
     * 状态
     *
     * Generated from protobuf field <code>int64 status = 7;</code>
     * @return int|string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * 状态
     *
     * Generated from protobuf field <code>int64 status = 7;</code>
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
     * 属性列表
     *
     * Generated from protobuf field <code>repeated .product.AttrItem items = 8;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getItems()
    {
        return $this->items;
    }

    /**
     * 属性列表
     *
     * Generated from protobuf field <code>repeated .product.AttrItem items = 8;</code>
     * @param array<\Product\AttrItem>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setItems($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Product\AttrItem::class);
        $this->items = $arr;

        return $this;
    }

}

