# encoding: utf-8
# copyright: 2017, Chef Software, Inc.
# license: All rights reserved
title 'Automate 2.0 Smoke Tests'

# Perform some basic requests against the backend services
{
  'authn-service': {
    path: '/api/v0/auth/tokens',
    response_match: {
    },
  },
  'authz-service': {
    path: '/api/v0/auth/policies/version',
    response_match: {
      name: 'authz-service'
    },
  },
  'config-mgmt-service': {
    path: '/api/v0/cfgmgmt/version',
    response_match: {
      name: 'config-mgmt-service'
    },
  },
  'compliance-service': {
    path: '/api/v0/compliance/reporting/version',
    response_match: {
      name: 'compliance'
    },
  },
  'ingest-service': {
    path: '/api/v0/ingest/version',
    response_match: {
      name: 'ingest-service'
    },
  },
#  'license-control-service': {
#    path: '/api/v0/license/status',
#    expected_status: 200,
#    response_match: {},
#  },
  'local-user-service': {
    path: '/api/v0/auth/users',
    response_match: {},
  },
  'notifications-service': {
    path: '/api/v0/notifications/version',
    response_match: {
      version: '1.0.0'
    },
  },
  'teams-service': {
    path: '/api/v0/auth/teams/version',
    response_match: {
      name: 'teams-service'
    },
  },
  'automate-gateway': {
    path: '/api/v0/gateway/version',
    response_match: {
      name: 'automate-gateway'
    },
  },
}.each_with_index do |(service_name, opts), index|
  control "automate-smoke-#{index}" do
    title "GET #{opts[:path]}"
    desc "Checks the version endpoint of #{service_name} to make sure it's up and running"

    describe automate_api_request(opts[:path]) do
      its('http_status') { should eq(opts[:expected_status] || 200) }
      its('parsed_response_body') { should include(opts[:response_match]) }
    end
  end
end
