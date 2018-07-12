<template>
    <div id="app">
        <el-row>
            <el-table
                :data="data"
                style="width: 100%">
                <el-table-column
                prop="Table"
                label="表名"
                width="180">
                </el-table-column>
                <el-table-column
                prop="Key"
                label="Key"
                width="180">
                </el-table-column>
                <el-table-column
                prop="Value"
                label="Value">
                </el-table-column>
            </el-table>
        </el-row>
        <el-row>
            <el-button @click="refresh">123467</el-button>
        </el-row>
    </div>
</template>

<script>
import Utils from '@/utils/ax'

export default {
  name: 'test',
  components: {
    Utils
  },
  data() {
    return {
      data: []
    }
  },
  mounted() {
    this.refresh()
  },
  methods: {
    refresh() {
      Utils.getDefault().then(row => {
        console.log(row)
        this.data = row.data
      })
    }
  }
}
</script>