import "./index.less"
import React, { Component } from 'react';
import { Card } from 'antd'
import ReactEcharts from 'echarts-for-react';
class Bar extends Component {
    constructor(props) {
        super(props)
        this.state = {
            sales: [5, 20, 36, 10, 10, 20],
            stores: [15, 120, 36, 110, 110, 20]
        }
    }
    /**
     * 柱状图的配置对象
     */
    getOption = (sales, stores) => {
        return {
            xAxis: {
                type: 'category',
                data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
            },
            yAxis: {
                type: 'value'
            },
            series: [{
                data: [150, 230, 224, 218, 135, 147, 260],
                type: 'line'
            }]
        }

    }
    getOption = (sales, stores) => {
        return {
            xAxis: {
                type: 'category',
                data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
            },
            yAxis: {
                type: 'value'
            },
            series: [{
                data: [150, 230, 224, 218, 135, 147, 260],
                type: 'line'
            }]
        }

    }
    setOption = () => {
        return {
            series: [
                {
                    name: '访问来源',
                    type: 'pie',
                    radius: '55%',
                    data: [
                        { value: 235, name: '视频广告' },
                        { value: 274, name: '联盟广告' },
                        { value: 310, name: '邮件营销' },
                        { value: 335, name: '直接访问' },
                        { value: 400, name: '搜索引擎' }
                    ]
                }
            ]
        }
    }

    render() {
        const { sales, stores } = this.state;
        return (
            <div className="home">
                <Card>
                    <ReactEcharts option={this.getOption(sales, stores)} />
                </Card>
                <Card>
                    <ReactEcharts option={this.setOption()} />
                </Card>
            </div>
        )
    }
}
export default Bar;