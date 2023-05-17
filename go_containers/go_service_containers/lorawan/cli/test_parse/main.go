package main

import (
        "encoding/json"
        "fmt"
)

type Payload struct {
        EndDeviceIds struct {
                DeviceId         string `json:"device_id"`
                ApplicationIds   struct {
                        ApplicationId string `json:"application_id"`
                } `json:"application_ids"`
                DevEui   string `json:"dev_eui"`
                DevAddr  string `json:"dev_addr"`
                Received string `json:"received_at"`
        } `json:"result"`
        UplinkMessage struct {
                FPort     int `json:"f_port"`
                FCnt      int `json:"f_cnt"`
                FrmPayload string `json:"frm_payload"`
                DecodedPayload struct {
                        Err int `json:"err"`
                        Messages []struct {
                                MeasurementId int `json:"measurementId"`
                                MeasurementValue float64 `json:"measurementValue"`
                                Type string `json:"type"`
                        } `json:"messages"`
                        Payload string `json:"payload"`
                        Valid bool `json:"valid"`
                } `json:"decoded_payload"`
                RxMetadata []struct {
                        GatewayIds struct {
                                GatewayId string `json:"gateway_id"`
                                Eui string `json:"eui"`
                        } `json:"gateway_ids"`
                        Timestamp int `json:"timestamp"`
                        Rssi int `json:"rssi"`
                        ChannelRssi int `json:"channel_rssi"`
                        Snr float64 `json:"snr"`
                        ChannelIndex int `json:"channel_index"`
                        Received string `json:"received_at"`
                } `json:"rx_metadata"`
                Settings struct {
                        DataRate struct {
                                Lora struct {
                                        Bandwidth int `json:"bandwidth"`
                                        SpreadingFactor int `json:"spreading_factor"`
                                        CodingRate string `json:"coding_rate"`
                                } `json:"lora"`
                        } `json:"data_rate"`
                        Frequency string `json:"frequency"`
                        Timestamp int `json:"timestamp"`
                } `json:"settings"`
                Received string `json:"received_at"`
                ConsumedAirtime string `json:"consumed_airtime"`
                VersionIds struct {
                        BrandId string `json:"brand_id"`
                        ModelId string `json:"model_id"`
                        HardwareVersion string `json:"hardware_version"`
                        FirmwareVersion string `json:"firmware_version"`
                        BandId string `json:"band_id"`
                } `json:"version_ids"`
                NetworkIds struct {
                        NetId string `json:"net_id"`
                        TenantId string `json:"tenant_id"`
                        ClusterId string `json:"cluster_id"`
                        ClusterAddress string `json:"cluster_address"`
                } `json:"network_ids"`
        } `json:"uplink_message"`
}


const jsonString = `{"result":{"end_device_ids":{"device_id":"lacima-test-seeed1","application_ids":{"application_id":"lacima-ranch-test-app-1"},"dev_eui":"2CF7F1203230FFFF","dev_addr":"260CB144"},"received_at":"2022-06-12T21:30:28.021492885Z","uplink_message":{"f_port":46,"f_cnt":197,"frm_payload":"I0VniQ==","decoded_payload":{},"rx_metadata":[{"gateway_ids":{"gateway_id":"lacima-ranch-1","eui":"58A0CBFFFE803E79"},"time":"2022-06-12T21:30:27.762223958Z","timestamp":2384327140,"rssi":-77,"channel_rssi":-77,"snr":9.25,"location":{"latitude":33.57851455733027,"longitude":-117.29939270420267,"altitude":731,"source":"SOURCE_REGISTRY"}},{"gateway_ids":{"gateway_id":"lacina-ranch-2","eui":"58A0CBFFFE803EA7"},"time":"2022-06-12T21:30:27.753279924Z","timestamp":2537042364,"rssi":-81,"channel_rssi":-81,"snr":9}],"settings":{"data_rate":{"lora":{"bandwidth":125000,"spreading_factor":7}},"coding_rate":"4/5","frequency":"905300000","timestamp":2384327140,"time":"2022-06-12T21:30:27.762223958Z"},"received_at":"2022-06-12T21:30:27.794033976Z","confirmed":true,"consumed_airtime":"0.051456s","version_ids":{"brand_id":"seeed","model_id":"loradevelopkit-e5","hardware_version":"1.0","firmware_version":"1.0","band_id":"US_902_928"},"network_ids":{"net_id":"000013","tenant_id":"ttn","cluster_id":"nam1","cluster_address":"nam1.cloud.thethings.network"}}}}`

func main() {
        var payload Payload

        err := json.Unmarshal([]byte(jsonString), &payload)
        if err != nil {
                fmt.Printf("Error: %v\n", err)
                return
        }

        fmt.Printf("Parsed JSON: %+v\n", payload)
}
