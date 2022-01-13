import { translate } from 'element-plus'
import { defineComponent, onMounted, reactive } from 'vue'
import * as api from "../../api/index"
import paginationVue from '../../components/pagination.vue'
export default defineComponent({
  components:{
    paginationVue
  },
  setup() {
    const state = reactive({
      articles: []
    })
    onMounted(async () => {
      let { data } = await api.getArticles({
        limit: 10,
        offset: 1
      })
      state.articles = data.result
      console.log(state.articles)
    })
    return () => (
      // CommentsCount: 0
      // Content: ""
      // Cover: ['']
      // CreatedAt: "2021-05-03T15:53:26Z"
      // DeletedAt: null
      // ID: 222
      // UpdatedAt: "2022-01-06T12:44:01.887Z"
      // article_path: "http://localhost:3800/markdown/ES6笔记/ES6函数语法糖.md"
      // hot: 0
      // like: 0
      // name: "ES6函数语法糖.md"
      // paragraph: ""
      // picture: ""
      // rec: 0
      // tag: "ES6笔记"
      // uuid: "85a05b1c-33d2-4a73-bd7b-1b2b0784f571"
      // whatch_number: 271
      <translate>
        <el-table data={state.articles} brother className="w-full" >
          <el-table-column prop="CreatedAt" label="创建时间" width="150"></el-table-column>
          <el-table-column prop="ID" label="文章ID" width="80"></el-table-column>
          <el-table-column prop="name" label="文章名字" width="80"></el-table-column>
          <el-table-column prop="like" label="点赞" width="80"></el-table-column>
          <el-table-column prop="tag" label="分类" width="80"></el-table-column>
          <el-table-column prop="whatch_number" label="查看次数" width="80"></el-table-column>
        </el-table>
      <paginationVue></paginationVue>
      </translate>
    )
  }
})

