import Vue from 'vue'
import { Button, FormModel, Input, Icon, message, Layout, Menu } from 'ant-design-vue'

message.config({
  top: '60px',
  duration: 2,
  maxCount: 3
})
Vue.prototype.$message = message

Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Menu)
