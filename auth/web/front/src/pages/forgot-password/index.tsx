import styles from './index.less';
import Layout from '../layout';
import { StepsForm, ProFormText } from '@ant-design/pro-form';
import { useIntl, FormattedMessage, Link } from 'umi';
import PhoneNumberForm from '@/components/PhoneNumberForm';
import { Button } from 'antd';

export default () => {
  const intl = useIntl();
  const onFinish = async (values: any) => {
    console.log(values);
  };
  return (
    <Layout
      title={intl.formatMessage({
        id: 'pages.forgotPassword',
      })}
    >
      <StepsForm
        onFinish={(values: any) => {
          console.log(values);
          return Promise.resolve(true);
        }}
        submitter={{
          render: (props) => {
            if (props.step === 0) {
              return (
                <Button type="primary" block onClick={() => props.onSubmit?.()}>
                  下一步
                </Button>
              );
            }
            return (
              <Button
                type="primary"
                key="goToTree"
                onClick={() => props.onSubmit?.()}
              >
                完成
              </Button>
            );
          },
        }}
      >
        <StepsForm.StepForm
          name="step1"
          title="工作信息"
          style={{ width: 330 }}
        >
          <PhoneNumberForm />
        </StepsForm.StepForm>
        <StepsForm.StepForm
          name="step2"
          title={'同步表单信息'}
          style={{ width: 330 }}
        >
          <ProFormText.Password
            label={intl.formatMessage({ id: 'pages.forgotPassword.new' })}
          />
          <ProFormText.Password
            label={intl.formatMessage({ id: 'pages.forgotPassword.confirm' })}
          />
        </StepsForm.StepForm>
      </StepsForm>
    </Layout>
  );
};
