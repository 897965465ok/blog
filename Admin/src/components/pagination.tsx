import { defineComponent, reactive, onMounted } from 'vue'
export default defineComponent({
    props: {
        message: {
            type: String, //设置数据类型
            default: "默认参数",
        },
    },
    setup(props) {
        return () => (
            <>
                <el-col class="w-full flex justify-center items-center">
                    <el-pagination layout="prev, pager, next" total={50}></el-pagination>
                </el-col >
            </>
        )
    }
})