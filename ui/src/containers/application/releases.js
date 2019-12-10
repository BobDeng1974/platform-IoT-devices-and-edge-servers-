import React, { useMemo } from 'react';
import moment from 'moment';
import { useNavigation } from 'react-navi';

import Card from '../../components/card';
import Table from '../../components/table';

const ReleasedBy = ({ release }) => {
  if (release) {
    if (release.createdByUser) {
      return `${release.createdByUser.firstName} ${release.createdByUser.lastName}`;
    } else if (release.createdByServiceAccount) {
      return release.createdByServiceAccount.name;
    }
  }
  return '-';
};

const Releases = ({
  route: {
    data: { params, application, releases },
  },
}) => {
  const navigation = useNavigation();
  const columns = useMemo(
    () => [
      { Header: 'Release ID', accessor: 'id' },
      {
        Header: 'Released by',
        Cell: ({ row: { original } }) => <ReleasedBy release={original} />,
      },
      {
        Header: 'Started',
        Cell: ({
          row: {
            original: { createdAt },
          },
        }) => moment(createdAt).fromNow(),
      },
      {
        Header: 'Device count',
        accessor: 'deviceCounts.allCount',
      },
    ],
    []
  );
  const tableData = useMemo(() => releases, [releases]);

  return (
    <Card
      title="Releases"
      size="large"
      actions={[
        {
          title: 'Create Release',
          href: `/${params.project}/applications/${application.name}/releases/create`,
        },
      ]}
    >
      <Table
        columns={columns}
        data={tableData}
        onRowSelect={({ id }) =>
          navigation.navigate(
            `/${params.project}/applications/${application.name}/releases/${id}`
          )
        }
      />
    </Card>
  );
};

export default Releases;
