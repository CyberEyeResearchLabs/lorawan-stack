// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { connect } from 'react-redux'

import {
  selectSelectedApplication,
  selectSelectedApplicationId,
  selectApplicationLinkIndicator,
  selectApplicationLinkFetching,
} from '../../store/selectors/applications'
import {
  selectCollaboratorsTotalCount,
  selectCollaboratorsFetching,
} from '../../store/selectors/collaborators'
import { selectApiKeysTotalCount, selectApiKeysFetching } from '../../store/selectors/api-keys'
import { selectDevicesTotalCount, selectDevicesFetching } from '../../store/selectors/devices'
import { getCollaboratorsList } from '../../store/actions/collaborators'
import { getApiKeysList } from '../../store/actions/api-keys'
import { getApplicationLink } from '../../store/actions/link'

const mapStateToProps = (state, props) => {
  const appId = selectSelectedApplicationId(state)
  const collaboratorsTotalCount = selectCollaboratorsTotalCount(state, { id: appId })
  const apiKeysTotalCount = selectApiKeysTotalCount(state)
  const devicesTotalCount = selectDevicesTotalCount(state)

  return {
    appId,
    application: selectSelectedApplication(state),
    collaboratorsTotalCount,
    apiKeysTotalCount,
    devicesTotalCount,
    link: selectApplicationLinkIndicator(state),
    statusBarFetching:
      collaboratorsTotalCount === undefined ||
      apiKeysTotalCount === undefined ||
      devicesTotalCount === undefined ||
      selectApplicationLinkFetching(state) ||
      selectDevicesFetching(state) ||
      selectApiKeysFetching(state) ||
      selectCollaboratorsFetching(state),
  }
}

const mapDispatchToProps = dispatch => ({
  loadData(appId) {
    dispatch(getCollaboratorsList('application', appId))
    dispatch(getApiKeysList('application', appId))
    dispatch(getApplicationLink(appId))
  },
})

export default ApplicationOverview =>
  connect(
    mapStateToProps,
    mapDispatchToProps,
  )(ApplicationOverview)
