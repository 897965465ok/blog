import React, { useState, useEffect } from "react";
import http from '../../http'
import "antd/dist/antd.css";
import "./index.less";
import {
  Button,
  Table,
  Input,
  Form,
  Row,
  Space
} from "antd";

const EditableTable = () => {
  // 创建from实例
  const [form] = Form.useForm();
  // 数据
  const [data, setData] = useState([]);
  // 是否被禁用
  const [editingKey, setEditingKey] = useState("");

  useEffect(async () => {
    let { data } = await http.getArticles()
    let list = data.result.map(item => {
      item.key = item.uuid
      return item
    })
    setData(list)
  }, [])


  // 判断那个单元格需要编辑
  const isEditing = (record) => {
    return record.key === editingKey

  }

  // 点击编辑触发
  const edit = (record) => {

    // 设置form-item k-v
    form.setFieldsValue({
      ...record
    });
    // 标记那个单元格需要需要修改
    setEditingKey(record.key);
  };

  // 取消
  const cancel = () => {
    setEditingKey("");
  };

  const save = async (key) => {

    try {

      // 当前需要修改的对象
      const row = await form.validateFields();
      // 拷贝到新data
      const newData = [...data];
      // 找到
      const index = newData.findIndex((item) => key === item.key);
      if (index > -1) {
        const item = newData[index];
        newData.splice(index, 1, { ...item, ...row });
        setData(newData);
        setEditingKey("");
      } else {
        newData.push(row);
        setData(newData);
        setEditingKey("");
      }

    } catch (errInfo) {
      console.log("Validate Failed:", errInfo);
    }
  };

  const columns = [
    {
      title: '文章',
      dataIndex: 'name',
      editable: true
    },
    {
      title: '分类',
      dataIndex: 'tag',
      editable: true
    },
    {
      title: '查看次数',
      dataIndex: 'whatch_number',
      editable: true
    }, {
      title: '赞',
      dataIndex: 'like',
      editable: true
    },
    {
      title: "编辑",
      render: (_, record) => {
        const editable = isEditing(record);
        return editable ? (
          <Space size="middle">
            <Button
              className="save-button"
              size="small"
              onClick={() => save(record.key)}>
              保存
            </Button>
            <Button
              className="cancel-button"
              size="small"
              onClick={() => setEditingKey("")}
            >取消</Button>
          </Space>
        ) : (
          <>
            <Button
              disabled={editingKey !== ""}
              onClick={() => edit(record)}
              size="small"
              type="primary"
            > 编辑</Button>
          </>
        )
      }
    }
  ];

  const Cell = ({
    editing,
    dataIndex,
    title,
    inputType,
    record,
    index,
    children,
    ...restProps
  }) => {
    // className: "ant-table-cell"
    // colSpan: null
    // rowSpan: null
    // style: {}

    return <td {...restProps}>
      {editing ? (
        <Form.Item
          name={dataIndex}
          style={{
            margin: 0
          }}
        >
          <Input size="small" />
        </Form.Item>
      ) : (children)}
    </td>

  }
  // 设置单元格
  const mergedColumns = columns.map((col) => {
    if (!col.editable) {
      return col;
    }
    return {
      ...col,
      onCell: (record) => ({
        record,
        inputType: "text",
        dataIndex: col.dataIndex,
        title: col.title,
        editing: isEditing(record)
      })
    };
  });

  return (
    <Form form={form} component={false}>
      <Table
        components={{
          body: {
            cell: Cell
          }
        }}
        dataSource={data}
        columns={mergedColumns}
        rowClassName="editable-row"
        pagination={{
          onChange: cancel
        }}
      />
    </Form>
  );
};

export default EditableTable