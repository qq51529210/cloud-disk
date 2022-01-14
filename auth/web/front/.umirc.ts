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
    {
      exact: true,
      path: '/forgot_password',
      title: 'Forgot Password',
      component: '@/pages/forgot-password',
    },
  ],
  fastRefresh: {},
  locale: {
    default: 'zh-CN',
    antd: true,
    baseNavigator: true,
  },
  mfsu: {},
});
