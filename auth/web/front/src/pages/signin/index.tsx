import styles from './index.less';
import Layout from '../layout';
import {
  LoginForm,
  ProFormText,
  ProFormCaptcha,
  ProFormSelect,
  ProFormCheckbox,
} from '@ant-design/pro-form';
import { useIntl, FormattedMessage, Link } from 'umi';
import {
  AlipayCircleOutlined,
  LockOutlined,
  MobileOutlined,
  TaobaoCircleOutlined,
  UserOutlined,
  WeiboCircleOutlined,
} from '@ant-design/icons';
import { Alert, Tabs, Space, Button, Divider, Row, Form } from 'antd';
import { useState, useEffect } from 'react';
import qrcode from 'qrcode';
import PhoneNumberForm from '@/components/PhoneNumberForm';

export default function IndexPage() {
  const intl = useIntl();
  const [type, setType] = useState<string>('account');
  const signinError = false;
  const onFinish = async (values: any) => {};
  const getVerificationCode = async (number: string) => {};
  useEffect(() => {
    if (type === 'qrcode') {
      let qrcodeContainer = document.getElementById('qrcode');
      if (qrcodeContainer) {
        qrcode.toCanvas(
          'test qrcode',
          { width: 256, errorCorrectionLevel: 'H' },
          (err: any, canvas: any) => {
            if (!err && canvas) {
              qrcodeContainer?.replaceChildren(canvas);
            }
          },
        );
      }
    }
  }, [type]);
  return (
    <Layout
      title={intl.formatMessage({
        id: 'pages.signin',
      })}
    >
      <Tabs size="large" type="card" centered onChange={setType}>
        <Tabs.TabPane
          key="account"
          tab={intl.formatMessage({
            id: 'pages.signin.tab.account',
          })}
        >
          <LoginForm
            subTitle=""
            onFinish={onFinish}
            submitter={{
              render: (props, doms) => (
                <>
                  <Button
                    block
                    type="primary"
                    onClick={() => props.form?.submit()}
                  >
                    {intl.formatMessage({
                      id: 'pages.signin',
                    })}
                  </Button>
                  <Row justify="end" align="middle" style={{ marginTop: 12 }}>
                    <Link to="/signup">
                      <FormattedMessage id="pages.signin.signup" />
                    </Link>
                    <Divider type="vertical" />
                    <Link to="/forgot_password">
                      <FormattedMessage id="pages.signin.forgotPassword" />
                    </Link>
                  </Row>
                </>
              ),
            }}
          >
            {signinError && (
              <Alert
                style={{
                  marginBottom: 24,
                }}
                message={intl.formatMessage({
                  id: 'pages.signin.errorMessage',
                })}
                type="error"
                showIcon
              />
            )}
            <ProFormText
              name="account"
              fieldProps={{
                size: 'large',
                prefix: <UserOutlined className={styles.prefixIcon} />,
              }}
              placeholder={intl.formatMessage({
                id: 'pages.signin.placeholder.account',
              })}
              rules={[
                {
                  required: true,
                  message: intl.formatMessage({
                    id: 'pages.signin.required.account',
                  }),
                },
              ]}
            />
            <ProFormText.Password
              name="password"
              fieldProps={{
                size: 'large',
                prefix: <LockOutlined className={styles.prefixIcon} />,
              }}
              placeholder={intl.formatMessage({
                id: 'pages.signin.placeholder.password',
              })}
              rules={[
                {
                  required: true,
                  message: intl.formatMessage({
                    id: 'pages.signin.required.password',
                  }),
                },
              ]}
            />
          </LoginForm>
        </Tabs.TabPane>
        <Tabs.TabPane
          key="phone"
          tab={intl.formatMessage({
            id: 'pages.signin.tab.phone',
          })}
        >
          <Form onFinish={onFinish}>
            <PhoneNumberForm />
            <Form.Item>
              <Button block type="primary" htmlType="submit" size="large">
                {intl.formatMessage({
                  id: 'pages.signup',
                })}
              </Button>
            </Form.Item>
          </Form>
        </Tabs.TabPane>
        <Tabs.TabPane
          key="qrcode"
          tab={intl.formatMessage({
            id: 'pages.signin.tab.qrcode',
          })}
        >
          <Row id="qrcode" justify="center"></Row>
        </Tabs.TabPane>
      </Tabs>
    </Layout>
  );
}
