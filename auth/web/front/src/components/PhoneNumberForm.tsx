import {
  ProFormCaptcha,
  ProFormSelect,
  ProFormText,
} from '@ant-design/pro-form';
import { MessageOutlined, MobileOutlined } from '@ant-design/icons';
import styles from './PhoneNumberForm.less';
import { useIntl } from 'umi';
import { Row, Col } from 'antd';

interface Props {
  width?: number;
}

export default (props: Props) => {
  const intl = useIntl();
  const getVerifyCode = async (number: string) => {};
  return (
    <>
      <Row gutter={80} style={{ minWidth: props.width }}>
        <Col span={6}>
          <ProFormSelect
            name="from"
            fieldProps={{
              size: 'large',
              allowClear: false,
            }}
            initialValue={'+86'}
            valueEnum={{ '+86': '+86' }}
          />
        </Col>
        <Col span={18}>
          <ProFormText
            name="number"
            fieldProps={{
              size: 'large',
              prefix: <MobileOutlined className={styles.prefixIcon} />,
            }}
            placeholder={intl.formatMessage({
              id: 'components.phoneNumber.placeholder',
            })}
            rules={[
              {
                required: true,
                message: intl.formatMessage({
                  id: 'components.phoneNumber.required',
                }),
              },
            ]}
          />
        </Col>
      </Row>
      <ProFormCaptcha
        name="captcha"
        fieldProps={{
          size: 'large',
          prefix: <MessageOutlined className={styles.prefixIcon} />,
        }}
        captchaProps={{
          size: 'large',
        }}
        placeholder={intl.formatMessage({
          id: 'components.verifyCode.placeholder',
        })}
        captchaTextRender={(timing, count) => {
          if (timing) {
            return `${count} ${intl.formatMessage({
              id: 'components.verifyCode.text',
            })}`;
          }
          return intl.formatMessage({
            id: 'components.verifyCode.text',
          });
        }}
        rules={[
          {
            required: true,
            message: intl.formatMessage({
              id: 'components.verifyCode.required',
            }),
          },
          // {
          //   pattern: /^[0-9]{6}$/,
          //   message: intl.formatMessage({
          //     id: 'components.verifyCode.error',
          //   }),
          // },
        ]}
        onGetCaptcha={getVerifyCode}
        allowClear={true}
      />
    </>
  );
};
