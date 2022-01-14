import styles from './layout.less';
import { SelectLang } from 'umi';
import { Typography, Divider, Row } from 'antd';

export default (props: any) => {
  return (
    <>
      <div className={styles.lang}>
        <SelectLang />
      </div>
      <Divider>
        <Typography>
          <Typography.Title>{props.title}</Typography.Title>
        </Typography>
      </Divider>
      <Row justify="center">
        <div style={{ width: 330 }}>{props.children}</div>
      </Row>
    </>
  );
};
