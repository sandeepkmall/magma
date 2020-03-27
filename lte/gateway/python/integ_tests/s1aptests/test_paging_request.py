"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import unittest


import gpp_types
import s1ap_types
import s1ap_wrapper
import time


class TestPagingRequest(unittest.TestCase):
    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()

    def tearDown(self):
        self._s1ap_wrapper.cleanup()

    def test_paging_request(self):
        """ Multi Enb Multi UE attach detach """
        # column is a enb parameter,  row is a number of enbs
        # column description: 1.Cell Id, 2.Tac, 3.EnbType, 4.PLMN Id
        enb_list = [(1, 1, 1, "001010")]

        self._s1ap_wrapper.multiEnbConfig(len(enb_list), enb_list)

        time.sleep(2)

        """ Attach, paging request """
        # Ground work.
        self._s1ap_wrapper.configUEDevice(1)
        req = self._s1ap_wrapper.ue_req
        ue_id = req.ue_id
        print(
            "************************* Running End to End attach for UE id ",
            ue_id,
        )
        # Now actually complete the attach
        self._s1ap_wrapper.s1_util.attach(
            ue_id,
            s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
            s1ap_types.tfwCmd.UE_ATTACH_ACCEPT_IND,
            s1ap_types.ueAttachAccept_t,
        )

        # Wait on EMM Information from MME
        self._s1ap_wrapper._s1_util.receive_emm_info()

        # Delay to ensure S1APTester sends attach complete before sending UE
        # context release
        time.sleep(0.5)

        print(
            "************************* Sending UE context release request ",
            "for UE id ",
            ue_id,
        )
        # Send UE context release request to move UE to idle mode
        ue_cntxt_rel_req = s1ap_types.ueCntxtRelReq_t()
        ue_cntxt_rel_req.ue_Id = ue_id
        ue_cntxt_rel_req.cause.causeVal = (
            gpp_types.CauseRadioNetwork.USER_INACTIVITY.value
        )
        self._s1ap_wrapper.s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_CNTXT_REL_REQUEST, ue_cntxt_rel_req
        )
        response = self._s1ap_wrapper.s1_util.get_response()
        self.assertEqual(
            response.msg_type, s1ap_types.tfwCmd.UE_CTX_REL_IND.value
        )

        time.sleep(0.3)
        print(
            "************************* Running UE downlink (UDP) for UE id ",
            ue_id,
        )
        with self._s1ap_wrapper.configDownlinkTest(
            req, duration=1, is_udp=True
        ) as test:
            response = self._s1ap_wrapper.s1_util.get_response()
            self.assertTrue(response, s1ap_types.tfwCmd.UE_PAGING_IND.value)
            # Send service request to reconnect UE
            ser_req = s1ap_types.ueserviceReq_t()
            ser_req.ue_Id = ue_id
            ser_req.ueMtmsi = s1ap_types.ueMtmsi_t()
            ser_req.ueMtmsi.pres = False
            ser_req.rrcCause = s1ap_types.Rrc_Cause.TFW_MO_DATA.value
            self._s1ap_wrapper.s1_util.issue_cmd(
                s1ap_types.tfwCmd.UE_SERVICE_REQUEST, ser_req
            )
            response = self._s1ap_wrapper.s1_util.get_response()
            self.assertEqual(
                response.msg_type, s1ap_types.tfwCmd.INT_CTX_SETUP_IND.value
            )
            test.verify()

        time.sleep(0.5)
        # Now detach the UE
        self._s1ap_wrapper.s1_util.detach(
            ue_id, s1ap_types.ueDetachType_t.UE_NORMAL_DETACH.value, True
        )
        time.sleep(0.5)


if __name__ == "__main__":
    unittest.main()