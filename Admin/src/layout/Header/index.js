
import { useState, createElement } from "react";
import { Menu, Layout } from "antd";
import { Link } from 'react-router-dom'
import * as Icon from '@ant-design/icons';
const Header = () => {
    let [style] = useState({
        background: "#ffffff"
    });
    let [menu] = useState([
        {
            title: "LOGO",
            uuid: Math.ceil(Math.random() * 1000),
            link: "/",
            iconName: "",
        }
    ]);
    return (
        <Layout.Header style={style}>
            <Menu mode="horizontal">
                {menu.map((item) => {
                    const { uuid, title, iconName, link } = item;
                    return (
                        <Menu.Item key="uuid">
                            <Link to={link}>{title}</Link>
                        </Menu.Item>
                    )
                })}
            </Menu>
        </Layout.Header>
    )
};
export default Header