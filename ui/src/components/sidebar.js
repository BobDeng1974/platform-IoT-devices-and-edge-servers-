import React from 'react';
import styled from 'styled-components';
import { useActive, useCurrentRoute } from 'react-navi';
import { Icon } from 'evergreen-ui';

import Logo from './icons/logo';
import { Row, Column, Link, Text } from './core';

const links = [
  {
    title: 'Devices',
    icon: 'multi-select',
    to: '/devices',
  },
  {
    title: 'Provisioning',
    icon: 'projects',
    to: '/provisioning',
  },
  {
    title: 'Applications',
    icon: 'applications',
    to: '/applications',
  },
  {
    title: 'IAM',
    icon: 'people',
    to: '/iam',
  },
  {
    title: 'Settings',
    icon: 'settings',
    to: '/settings',
  },
];

const SidebarLink = styled(Link)`
  display: flex;
  justify-content: center;
  align-items: center;
  transition: background-color 200ms;
  border-radius: 4px;
  text-transform: uppercase;
  text-decoration: none !important;

  background-color: ${props =>
    props.active ? props.theme.colors.grays[0] : 'inherit'};
  color: ${props =>
    props.active ? props.theme.colors.primary : props.theme.colors.white};

  &:hover {
    background-color: ${props => props.theme.colors.grays[0]};
  }

  & span {
    color: ${props =>
      props.active ? props.theme.colors.primary : props.theme.colors.white};
  }

  & > div > svg {
    fill: ${props =>
      props.active
        ? props.theme.colors.primary
        : props.theme.colors.white} !important;
  }

  &:last-child {
    margin-top: auto;
  }
`;

const Sidebar = () => {
  const route = useCurrentRoute();

  if (!route) {
    return null;
  }

  const projectSelected = !!route.data.params.project;

  return (
    <Column
      width={136}
      alignSelf="stretch"
      bg="black"
      alignItems="center"
      flexShrink={0}
    >
      <Row paddingY={5}>
        <Link href="/">
          <Logo />
        </Link>
      </Row>
      {projectSelected &&
        links.map(({ to, title, icon }) => {
          const href = `/${route.data.params.project}${to}`;

          return (
            <SidebarLink
              href={href}
              width={120}
              paddingY={4}
              marginBottom={2}
              key={title}
              fontSize={0}
              active={useActive(href, { exact: false })}
            >
              <Column alignItems="center">
                <Icon icon={icon} color="white" size={24} />
                <Text marginTop={3}>{title}</Text>
              </Column>
            </SidebarLink>
          );
        })}
    </Column>
  );
};

export default Sidebar;
