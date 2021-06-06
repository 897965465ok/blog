import { Menu } from "antd";
import { useState, createElement } from "react";
import { Link } from 'react-router-dom'
import * as Icon from '@ant-design/icons';
const Left = () => {
    let [style] = useState({
        width: 300,
    });
    let [menu] = useState([
        {
            title: "首页管理",
            uuid: Math.ceil(Math.random() * 1000),
            iconName: "HomeOutlined",
            chlidren: [
                {
                    title: "文章",
                    uuid: Math.ceil(Math.random() * 1000),
                    link: "/article",
                    chlidren: [],
                },
                {
                    title: "轮播图",
                    uuid: Math.ceil(Math.random() * 1000),
                    link: "/banner",
                    chlidren: [],
                },
                {
                    title: "推荐文章",
                    uuid: Math.ceil(Math.random() * 1000),
                    link: "/",
                    chlidren: [],
                },
            ]
        }
    ]);
    return (
        <Menu
            style={style}
            mode="inline"
        >
            {menu.map((item) => {
                let { uuid, title, chlidren, iconName } = item;
                return (
                    <Menu.SubMenu key={uuid} title={title} icon={createElement(Icon[iconName], {
                        style: {
                            fontSize: "20px"
                        }
                    })} >
                        {chlidren.map((item) => {
                            return <Menu.Item key={item.uuid}>
                                <Link to={item.link}>{item.title}</Link>
                            </Menu.Item>;
                        })}
                    </Menu.SubMenu>
                );
            })}
        </Menu>
    );
};

export default Left;
