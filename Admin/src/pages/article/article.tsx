import { defineComponent, onMounted, reactive } from 'vue'
import * as api from "../../api/index"
export default defineComponent({
  setup() {
    const state = reactive({
      articles: []
    })
    onMounted(async () => {
      let { data } = await api.getArticles()
      state.articles = data.result
      console.log(state.articles)
    })
    return () => (
      <el-table data={state.articles} className=" w-full " >
        <el-table-column prop="CreatedAt" label="创建时间" width="150"></el-table-column>
        <el-table-column prop="ID" label="文章ID" width="80"></el-table-column>
      </el-table>
    )
  }
})

