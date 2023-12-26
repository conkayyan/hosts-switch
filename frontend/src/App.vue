<template>
  <el-row class="tac">
    <el-col :span="8" :lg="4" class="border-right">
      <el-menu
          class="no-border-right"
          default-active="all_hosts"
          @select="handleSelect"
          @open="handleSelect"
          @close="handleSelect"
      >
        <el-menu-item index="all_hosts">
          <el-icon>
            <document/>
          </el-icon>
          <span>All Hosts</span>
        </el-menu-item>
        <el-menu-item index="in_use">
          <el-icon>
            <el-icon-document-checked/>
          </el-icon>
          <span>In Use</span>
        </el-menu-item>
        <el-sub-menu index="host_groups">
          <template #title>
            <el-icon>
              <icon-menu/>
            </el-icon>
            <span>Host Groups</span>
          </template>
          <el-sub-menu :index="'show_group:'+groupName" v-for="(group, groupName) in listByGroup.list">
            <template #title>{{groupName}}<el-col class="menu-switch"><el-switch v-model="group.show" @change="handleSwitchByGroupName(group)" @click.stop /></el-col></template>
            <el-menu-item :index="'group_'+id" :title="row.ip" v-for="(row, id) in group.list"><el-checkbox v-model="row.show" @change="handleSwitchByHostnameId(row)" :label="row.hostname" size="large" /></el-menu-item>
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
    <el-col :span="16" class="show-content">
      <el-tabs v-model="activeName" class="ml-2 mr-2" @tab-click="handleClickTab">
        <el-tab-pane label="Normal" name="normal">
          <el-form v-if="activeIndex==='all_hosts'" :model="allHostsForm">
            <el-form-item>
              <el-alert class="el-form-item__content" effect="dark"
                        title="e.g. 127.0.0.1 www.domain.com # Group Name One # Group Name Two" type="info"/>
            </el-form-item>
            <el-form-item class="mt-2">
              <CodeEditor v-model="allHostsForm.allHostsText"/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmitAllHosts">Save</el-button>
              <el-button @click="copyToClipboard">Copy to Clipboard</el-button>
            </el-form-item>
          </el-form>
          <el-form v-else-if="activeIndex==='add_host'" :model="addHostForm" label-width="120px">
            <el-form-item label="Group Name">
              <el-input v-model="addHostForm.groupName" class="el-col-10" placeholder="Group Name"/>
            </el-form-item>
            <el-form-item label="IP">
              <el-input v-model="addHostForm.ip" class="el-col-10" placeholder="e.g. 127.0.0.1"/>
            </el-form-item>
            <el-form-item label="Host">
              <el-input v-model="addHostForm.hostname" class="el-col-10" placeholder="e.g. www.domain.com"/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmitAddHost">Add</el-button>
            </el-form-item>
          </el-form>
          <el-table v-else-if="activeIndex==='show_group'"
                    :key="groupName"
                    :data="tableData"
                    stripe
                    style="width: 100%"
          >
            <el-table-column align="center" label="Active" width="80">
              <template #default="scope">
                <el-checkbox v-model="scope.row.show" size="large" @change="handleSwitchByHostnameId(scope.row)"/>
              </template>
            </el-table-column>
            <el-table-column label="IP" prop="ip" width="150"/>
            <el-table-column label="Hostname" prop="hostname"/>
            <el-table-column label="Operations" width="200">
              <template #default="scope">
                <!--            <el-button size="small" @click="handleEditHost(scope.$index, scope.row)">Edit</el-button>-->
                <el-button size="small" type="danger" @click="handleDeleteHost(scope.$index, scope.row)">Delete
                </el-button>
              </template>
            </el-table-column>
          </el-table>
          <el-table v-else-if="activeIndex==='in_use'"
                    :key="groupName"
                    :data="tableData"
                    stripe
                    style="width: 100%"
          >
            <el-table-column align="center" label="Active" width="80">
              <template #default="scope">
                <el-checkbox v-model="scope.row.show" size="large" @change="handleSwitchByHostnameId(scope.row)"/>
              </template>
            </el-table-column>
            <el-table-column label="IP" prop="ip" width="150"/>
            <el-table-column label="Hostname" prop="hostname"/>
            <el-table-column label="Group Name" prop="group_name"/>
            <el-table-column label="Operations" width="150">
              <template #default="scope">
                <!--            <el-button size="small" @click="handleEditHost(scope.$index, scope.row)">Edit</el-button>-->
                <el-button size="small" type="danger" @click="handleDeleteHost(scope.$index, scope.row)">Delete
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
        <el-tab-pane label="Advanced" name="advanced">
          <el-form v-if="activeIndex==='all_hosts'" :model="allHostsForm">
            <el-form-item>
              <el-alert class="el-form-item__content" effect="dark"
                        title="e.g. 127.0.0.1 www.domain.com # Group Name One # Group Name Two" type="info"/>
            </el-form-item>
            <el-form-item class="mt-2">
              <CodeEditor v-model="allHostsForm.allHostsText"/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmitAllHosts">Save</el-button>
              <el-button @click="copyToClipboard">Copy to Clipboard</el-button>
            </el-form-item>
          </el-form>
          <el-form v-else-if="activeIndex==='add_host'" :model="addHostForm" label-width="120px">
            <el-form-item label="Group Name">
              <el-input v-model="addHostForm.groupName" class="el-col-10" placeholder="Group Name"/>
            </el-form-item>
            <el-form-item label="IP">
              <el-input v-model="addHostForm.ip" class="el-col-10" placeholder="e.g. 127.0.0.1"/>
            </el-form-item>
            <el-form-item label="Host">
              <el-input v-model="addHostForm.hostname" class="el-col-10" placeholder="e.g. www.domain.com"/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmitAddHost">Add</el-button>
            </el-form-item>
          </el-form>
          <el-table v-else-if="activeIndex==='show_group'"
                    :key="groupName"
                    :data="tableData"
                    stripe
                    style="width: 100%"
          >
            <el-table-column align="center" label="Active" width="80">
              <template #default="scope">
                <el-checkbox v-model="scope.row.show" size="large" @change="handleSwitchByHostnameId(scope.row)"/>
              </template>
            </el-table-column>
            <el-table-column label="IP" prop="ip" width="150"/>
            <el-table-column label="Hostname" prop="hostname"/>
            <el-table-column label="Operations" width="200">
              <template #default="scope">
                <!--            <el-button size="small" @click="handleEditHost(scope.$index, scope.row)">Edit</el-button>-->
                <el-button size="small" type="danger" @click="handleDeleteHost(scope.$index, scope.row)">Delete
                </el-button>
              </template>
            </el-table-column>
          </el-table>
          <el-table v-else-if="activeIndex==='in_use'"
                    :key="groupName"
                    :data="tableData"
                    stripe
                    style="width: 100%"
          >
            <el-table-column align="center" label="Active" width="80">
              <template #default="scope">
                <el-checkbox v-model="scope.row.show" size="large" @change="handleSwitchByHostnameId(scope.row)"/>
              </template>
            </el-table-column>
            <el-table-column label="IP" prop="ip" width="150"/>
            <el-table-column label="Hostname" prop="hostname"/>
            <el-table-column label="Group Name" prop="group_name"/>
            <el-table-column label="Operations" width="150">
              <template #default="scope">
                <!--            <el-button size="small" @click="handleEditHost(scope.$index, scope.row)">Edit</el-button>-->
                <el-button size="small" type="danger" @click="handleDeleteHost(scope.$index, scope.row)">Delete
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {Document, DocumentAdd, Menu as IconMenu} from '@element-plus/icons-vue'
import CodeEditor from './components/CodeEditor.vue'
import {reactive, ref} from "vue"
import type {TabsPaneContext} from 'element-plus'
import {ElMessage} from 'element-plus'
import 'element-plus/dist/index.css'
import {
  AddHost,
  DeleteHost,
  GetHostsText,
  GetList,
  GetListByGroup,
  SaveAllHosts,
  SwitchByGroupName,
  SwitchByHostnameId
} from "../wailsjs/go/main/App"
import {ClipboardSetText} from "../wailsjs/runtime"

const activeName = ref('normal')
const groupName = ref('')
const activeIndex = ref('all_hosts')
const listByGroup = reactive({list:{}})
const addHostForm = reactive({
  groupName: '',
  ip: '',
  hostname: '',
})
const allHostsForm = reactive({
  allHostsText: ''
})
const tableData = ref([])

function getHostsText() {
  GetHostsText().then(result => {
    allHostsForm.allHostsText = result
  })
}
getHostsText()

function getListByGroup() {
  GetListByGroup().then(result => {
    console.log("listByGroup", result)
    listByGroup.list = result

    if (groupName.value !== "" && activeIndex.value === "show_group") {
      tableData.value.splice(0, tableData.value.length)
      let groupInfo = listByGroup.list[groupName.value].list
      for (let k in groupInfo){
        tableData.value.push(groupInfo[k])
      }
      console.log(tableData)
    }
  })
}
getListByGroup()

function getList() {
  GetList().then(result => {
    console.log("list", result)

    tableData.value.splice(0, tableData.value.length)
    for (let k in result){
      if (result[k].show){
        tableData.value.push(result[k])
      }
    }
  })
}

const handleSelect = (key: string, keyPath: string[]) => {
  console.log('select', key, keyPath)
  if (key === 'all_hosts') {
    activeIndex.value = key
  } else if (key === 'add_host') {
    activeIndex.value = key
  } else if (key === 'in_use') {
    activeIndex.value = key
    groupName.value = key
    getList()
  } else if (key.substring(0, 10) === 'show_group') {
    activeIndex.value = key.substring(0, 10)
    tableData.value.splice(0, tableData.value.length)
    groupName.value = key.substring(11)
    let groupInfo = listByGroup.list[groupName.value].list
    for (let k in groupInfo){
      tableData.value.push(groupInfo[k])
    }
    console.log(tableData)
  }
}

const handleSwitchByGroupName = (group) => {
  console.log('switch', group.group_name, group.show)
  SwitchByGroupName(group.group_name, group.show).then(result => {
    if (result!==''){
      ElMessage.error('switch failed!' + result)
    }else{
      getHostsText()
      getListByGroup()
      ElMessage.success('switch successfully!')
    }
  })
}

const handleSwitchByHostnameId = (row) => {
  console.log('switch', row.group_name, row.id, row.hostname, row.show)
  SwitchByHostnameId(row.group_name, row.id, row.show).then(result => {
    if (result!==''){
      ElMessage.error('switch failed!' + result)
    }else{
      getHostsText()
      getListByGroup()
      ElMessage.success('switch successfully!')
    }
  })
}

const handleDeleteHost = (index, row) => {
  console.log('delete', index, row.group_name, row.hostname, row.id)
  DeleteHost(row.group_name, row.id).then(result => {
    if (result!==''){
      ElMessage.error('delete failed!' + result)
    }else{
      tableData.value.splice(index, 1)
      getHostsText()
      getListByGroup()
      ElMessage.success('delete successfully!')
    }
  })
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
      getHostsText()
      getListByGroup()
      ElMessage.success('save successfully!')
    }
  })
}

const onSubmitAllHosts = () => {
  console.log('submit all hosts!')
  SaveAllHosts(allHostsForm.allHostsText).then(result => {
    if (result!==''){
      ElMessage.error('save failed!' + result)
    }else{
      getHostsText()
      getListByGroup()
      ElMessage.success('save successfully!')
    }
  })
}

const copyToClipboard = () => {
  console.log('copy to clipboard!')
  ClipboardSetText(allHostsForm.allHostsText).then(result => {
    if (!result){
      ElMessage.error('copy to clipboard failed!')
    }else{
      ElMessage.success('copy to clipboard successfully!')
    }
  })
}

const handleClickTab = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event)
}
</script>

<style scoped>
.show-content{height: 100vh;}
.menu-switch{position: absolute; right: 0; padding-right: 50px;}
.ml-2 {margin-left: 20px;}
.mr-2 {margin-right: 20px;}
.mt-2 {margin-top: 20px;}
.border-right {border-right: 1px solid var(--el-border-color);}
.no-border-right { border: none; }
</style>