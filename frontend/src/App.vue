<template>
  <el-row class="tac">
    <el-col :span="8" class="border-right">
      <el-menu
          class="no-border-right"
          default-active="all_hosts"
          @select="handleSelect"
      >
        <el-menu-item index="all_hosts">
          <el-icon>
            <document/>
          </el-icon>
          <span>All Hosts</span>
        </el-menu-item>
        <el-sub-menu index="host_groups">
          <template #title>
            <el-icon>
              <icon-menu/>
            </el-icon>
            <span>Host Groups</span>
          </template>
          <el-sub-menu :index="groupName" v-for="(group, groupName) in hostsList.list">
            <template #title>{{groupName}}<el-switch class="ml-2" v-model="group.show" @click.stop /></template>
            <el-menu-item :index="row.hostname" :title="row.ip" v-for="row in group.list"><el-checkbox v-model="row.show" :label="row.hostname" size="large" /></el-menu-item>
          </el-sub-menu>
        </el-sub-menu>
        <el-menu-item index="add_host">
          <el-icon>
            <document-add/>
          </el-icon>
          <span>Add Host</span>
        </el-menu-item>
      </el-menu>
    </el-col>
    <el-col :span="16">
      <el-form :model="allHostsForm" class="mt-2 ml-2" v-if="activeIndex==='all_hosts'">
        <el-form-item>
          <CodeEditor v-model="allHostsForm.allHosts" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmitAllHosts">Save</el-button>
        </el-form-item>
      </el-form>
      <el-form :model="addHostForm" label-width="120px" class="mt-2" v-else-if="activeIndex==='add_host'">
        <el-form-item label="Group Name">
          <el-input v-model="addHostForm.groupName" placeholder="Group Name" class="el-col-10"/>
        </el-form-item>
        <el-form-item label="IP">
          <el-input v-model="addHostForm.ip" placeholder="e.g. 127.0.0.1" class="el-col-10"/>
        </el-form-item>
        <el-form-item label="Host">
          <el-input v-model="addHostForm.hostname" placeholder="e.g. www.domain.com" class="el-col-10"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmitAddHost">Add</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import { DocumentAdd, Document, Menu as IconMenu, Location, Setting } from '@element-plus/icons-vue'
import CodeEditor from './components/CodeEditor.vue'
import {reactive, ref} from "vue";
import {ElMessage} from 'element-plus'
import 'element-plus/dist/index.css'
import {AddHost, GetHosts, GetHostsList, SaveAllHosts} from "../wailsjs/go/main/App";

const activeIndex = ref('all_hosts')
const hostsList = reactive({list:{}})
const addHostForm = reactive({
  groupName: '',
  ip: '',
  hostname: '',
})
const allHostsForm = reactive({
  allHosts: ''
})

function getHosts() {
  GetHosts().then(result => {
    allHostsForm.allHosts = result
  })
}
getHosts()

function getHostsList() {
  GetHostsList().then(result => {
    console.log("list", result)
    hostsList.list = result
  })
}
getHostsList()

const handleSelect = (key: string, keyPath: string[]) => {
  console.log('select', key, keyPath)
  if (key === 'all_hosts') {
    activeIndex.value = key
  } else if (key === 'add_host') {
    activeIndex.value = key
  }
}

const onSubmitAddHost = () => {
  console.log('submit add host!')
  AddHost(addHostForm.groupName, addHostForm.ip, addHostForm.hostname).then(result => {
    if (result!==''){
      ElMessage.error('save failed!' + result)
    }else{
      addHostForm.groupName = ''
      addHostForm.ip = ''
      addHostForm.hostname = ''
      ElMessage.success('save successfully!')
      getHosts()
      getHostsList()
    }
  })
}

const onSubmitAllHosts = () => {
  console.log('submit all hosts!')
  SaveAllHosts(allHostsForm.allHosts).then(result => {
    if (result!==''){
      ElMessage.error('save failed!' + result)
    }else{
      ElMessage.success('save successfully!')
      getHosts()
      getHostsList()
    }
  })
}
</script>

<style scoped>
.ml-2 {margin-left: 20px;}
.mt-2 {margin-top: 20px;}
.border-right {border-right: 1px solid var(--el-border-color);}
.no-border-right { border: none; }
</style>