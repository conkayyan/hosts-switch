<template>
  <el-row class="tac">
    <el-col :span="6">
      <el-menu
          default-active="1-1-1"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
      >
        <el-sub-menu index="1">
          <template #title>
            <el-icon>
              <icon-menu/>
            </el-icon>
            <span>Host Groups</span>
          </template>
          <el-sub-menu index="1-1">
            <template #title>Group One<el-switch class="ml-2" v-model="value1" @click.stop /></template>
            <el-menu-item index="1-1-1"><el-checkbox v-model="checked1" label="item one" size="large" /></el-menu-item>
            <el-menu-item index="1-1-2"><el-checkbox v-model="checked2" label="item two" size="large" /></el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="1-2">
            <template #title>Group Two<el-switch class="ml-2" v-model="value2" @click.stop /></template>
            <el-menu-item index="1-2-1"><el-checkbox v-model="checked3" label="item one" size="large" /></el-menu-item>
            <el-menu-item index="1-2-2"><el-checkbox v-model="checked4" label="item two" size="large" /></el-menu-item>
          </el-sub-menu>
        </el-sub-menu>
        <el-menu-item index="2">
          <el-icon>
            <document/>
          </el-icon>
          <span>All Hosts</span>
        </el-menu-item>
      </el-menu>
    </el-col>
    <el-col :span="18">
      <el-form :model="form" label-width="120px">
        <el-form-item label="Group Name">
          <el-input v-model="form.groupName" placeholder="Group Name" />
        </el-form-item>
        <el-form-item label="Group Items">
          <CodeEditor v-model="form.groupItems" :default-value="hosts" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">Add</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'
import CodeEditor from './components/CodeEditor.vue'
import {reactive, ref} from "vue";

const value1 = ref(true)
const value2 = ref(true)
const checked1 = ref(true)
const checked2 = ref(true)
const checked3 = ref(true)
const checked4 = ref(true)
const form = reactive({
  groupName: '',
  groupItems: '',
})

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

const onSubmit = () => {
  console.log('submit!')
}
</script>

<style scoped>
.ml-2 {margin-left: 20px;}
</style>