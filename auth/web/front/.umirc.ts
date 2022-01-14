import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    { exact: true, path: '/', redirect: '/signin' },
    {
      exact: true,
      path: '/signin',
      title: 'Sign In',
      component: '@/pages/signin',
    },
    {
      exact: true,
      path: '/signup',
      title: 'Sign Up',
      component: '@/pages/signup',
    },
  ],
  fastRefresh: {},
  locale: {
    // default zh-CN
    default: 'zh-CN',
    antd: true,
    // default true, when it is true, will use `navigator.language` overwrite default
    baseNavigator: true,
  },
  musf: {},
});
