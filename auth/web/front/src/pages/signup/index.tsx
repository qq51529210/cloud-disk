import styles from './index.less';
import Layout from '../layout';
import { ProFormText } from '@ant-design/pro-form';
import { useIntl, FormattedMessage, Link } from 'umi';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Row, Button, Form } from 'antd';
import PhoneNumberForm from '@/components/PhoneNumberForm';

export default () => {
  const intl = useIntl();
  const onFinish = async (values: any) => {
    console.log(values);
  };
  return (
    <Layout
      title={intl.formatMessage({
        id: 'pages.signup',
      })}
    >
      <Form onFinish={onFinish}>
        <ProFormText
          name="account"
          fieldProps={{
            size: 'large',
            prefix: <UserOutlined className={styles.prefixIcon} />,
          }}
          placeholder={intl.formatMessage({
            id: 'pages.signup.placeholder.account',
          })}
          rules={[
            {
              required: true,
              message: intl.formatMessage({
                id: 'pages.signup.required.account',
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
            id: 'pages.signup.placeholder.password',
          })}
          rules={[
            {
              required: true,
              message: intl.formatMessage({
                id: 'pages.signup.required.password',
              }),
            },
          ]}
        />
        <PhoneNumberForm />
        <Form.Item>
          <Button block type="primary" htmlType="submit" size="large">
            {intl.formatMessage({
              id: 'pages.signup',
            })}
          </Button>
          <Row justify="end" align="middle" style={{ marginTop: 12 }}>
            <Link to="/signin">
              <FormattedMessage id="pages.signup.signin" />
            </Link>
          </Row>
        </Form.Item>
      </Form>
    </Layout>
  );
};
