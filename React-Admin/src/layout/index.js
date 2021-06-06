import "./index.less"
import Header from './Header'
import Left from "./Left"
import RouterGroup from "../router/router"
import { Layout } from 'antd';
import { BrowserRouter as Router } from "react-router-dom";
import { StoreContext } from 'redux-react-hook';
import { store } from '../store'
const { Footer, Content } = Layout;
export default () => (
    <StoreContext.Provider value={store}>
        <Router>
            <Layout>
                <Header></Header>
                <Content className="main">
                    <Left></Left>   
                    <RouterGroup />
                </Content>
                {/* <Footer></Footer> */}
            </Layout>
        </Router>
    </StoreContext.Provider>
);

