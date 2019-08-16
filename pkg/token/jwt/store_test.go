package jwt

import (
	"testing"
	"time"

	"elbix.dev/engine/pkg/sec"
	"github.com/stretchr/testify/require"
)

// privateKeyTest is a private RSA key for test (base64 encoded)
const privateKeyTest = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS3dJQkFBS0NBZ0VBbUlsSTZLK3Brb1dMa2k3czcvcVhae" +
	"HE3dWwwUUlqZVVBTFNPTGxiSUE5cmJQb2FSCmRIYmdWNHE1eXpsanhIY0YwL0Z1SE02WTV0WFJYVUF5WHlJRmw2bkN0cUdXU0pzRndJNk" +
	"N6VmNQc2h4ZzVYMzAKTDM4RFlpUk9zRVA2c08vV1oxK3lSYXRPeDNPelNSYTV4b0drTmp1bWkyd3E4a3JvR1dPb0FkdlFEa1RKNGJPOAp" +
	"PMjBCM1FlNmk5QmJmbkZiOUtMTFVJK0FiQmFBMlpyS0c3SlF6dmRsT0hrU0tSdDV4NS9HalhyN2Q5QlQxZ1pkCjMyM0UvYlExaWRCQXZz" +
	"ODBuMytJNmwwL1hWdnlRSUhwelc3ZEdNWnVEc3Z0bllNUlMwQXFhb0h6QnlKZE8xa3EKWHduVHBZMHhQRzFjeHJzNVdmVWVOVVpMa0pmN" +
	"ExmT3d4UjJ4ckozLzNMdmdCR09XOHdKQ3BLSm9Ha2w4cDVEWgpTMElnWFoyTTRhbjEvNzVuQnJRL1poTWUrVnpMSTFYLzFwR0dESndhVU" +
	"8rMDA5aHllZnlzQVFUMGxjOUNsZHU1CjllMnZybGZKUmtkSVhjTHdhZkRWNlNTSVVNYUJKVlNTcnFNZnRLUkp4bUgweG5TNVpuRWgxTnN" +
	"rV29JSzIzYWIKOXpyaEhESk5MTnhQeHJPMTFlRjJ6MmRzeEc3Y1ZpdnovVDVWNTB4QUN6RE9oNHdtZ0hOOXJ0UG1CbkxPSjVEcwpoU3dV" +
	"WWdLcTdkSE1nZDQ5R1Rqd0FFaUJrdVFUczladjNabUdWZVE3Z1NhaWpjcWllanl0ME9vUUh2ZnYreHlOCi8rOWRicE43UUNkS3dleDRnM" +
	"kZLSittM0tKeDR6eW5LUDZaaU9mL0psUlVQcmZVbUpHWUFUY3AwaTNFQ0F3RUEKQVFLQ0FnRUFpN25PR0orNDVLZ0NOdkZYUjJpMFFkVW" +
	"lPOCs5MUtvMm9BTEU3enU5VXBLckhQWmx6VnFMbklWTwpOaE9uQ1RQQ0dlaGhabS8xYmhaVkIyZWovTE1Nam1ZS0lBT2F4MzJCNG1abyt" +
	"USVVCYUtyTysySnZleEtMQjk0CkdVemNHakpJb3o3czdaYkR2ajZFNzZHcW9XVG1DR3ZSM290Q045cFlDOUhXOTRUa2EvMmxOR1BSUEZw" +
	"WlI2QXAKeUtNZS9JOVluWnAyUmFCK0VnVDNHTjRLM3g0VG1kTFRYTXpDenZjRk1kQnhMKzZFNWJibTVQbkdDY0hpbmtmeQpkd0JXSHJOd" +
	"HorRWlpUlRMMDNoRFM4WDFacUowVTVxSS9ITGZnZGFyQTlTT1Jpd3RjODZOdXF5dC9JcmkzaGFkCkVHNVpldzNMdGxTcnpEOU5QdkdHZk" +
	"lscTF0V2hoMFROTFNnamNGQ2RaRHkzclRHYXczanBGMlhMK21QVGErYkgKRVNkTXNwWHdRZXdDY09lV2p2RWtNYXgwbVcxNHljdk1MOGR" +
	"BbmtXdmxGUTdLcytneUtSMGNhKzNTS3VNanJTZwo4aGlJanNZZXZYMjBUR0RWT0dHNUx0YUVIQldTNzZhdUpzNW5YV2hMdWt2RDVmM2ll" +
	"cElNTXUwazRXVk5NRk5ICkp3NkFjM1NBODZmdTRLdDdxTHZ3UG1La1ZiZzJZdTIyd1dEK0ZPRUM1T2xSK2c3MjJZc0ZCU1h2S2puMm5TZ" +
	"24KK3NoaStGZXJ0SzVXOVZXa0Z4bmlxZ3BrNjFPM0p6dnZyQU5GSWc1OHRidEdQV2psSGtrR3Q4RjZZN29kZ29mZwp3VzV2MXdSVTdtMS" +
	"82Tlc2SUhid3RUaTlhMFVnZXpUS1FyVlRRd2dXNkVnQkJwdDd0cVVDZ2dFQkFNaXFadXVmCjlqRkxuZlVYbXd5ODhCL1JaK0w1S09Zc2l" +
	"PZGZjODg5MThHVExmUnQ0OHdJcFlzU05hblVjWk8wZkVvSkJObkMKcmFobXpob082V3pYMFM0ZklQcHJPS1BHQmlueUF5endvR2pZclRH" +
	"UktyR2ViejJOV3BVOFcrUzdJanVaWnNsMAovRmgwdkdwcWY3WnJrZCs5SlF4MnUwK3VqTnlMUUlrRi9CWnpPK3hXVDBlM2RFOFZEUmpUa" +
	"TZtZ3YzcjRhZlBMCm1UblMyQ25mL1NGZytuamN6UXZRZlVxS292Z3dldzVWZXk0TVkyeTN5d2xNTHNTRE5EK1h4UzRrNTVwN0R5TVUKOH" +
	"BteWZqWjJSU0tsV2h6ZG1QT1JBRmkxdk9ldXV1aFJGbmtwTjVCM1FjaHFLb3JrWG5qVndMZGR6T2JocG1FUApwQVgzV042b0FLWFp5T2N" +
	"DZ2dFQkFNS1pTTStOZHJnUk9STzM3bFdkU3VONHR4a0U3aVMybDVjYmxVNXFVTlpMCnVMdEw4a3J4eDZrcDhUYWlsMS9IRmMzYlM2L3hE" +
	"ejExSDY4RWV2OFRmZlhJNyswcG1xTTNGVWtDcGdRY0s0ZlQKUysycE5nUUFGSVBpUnFsTnZhZm5HaEV4eDl3NWNqQ3A5TklCQ0swekRHU" +
	"k9Pb2pkendMdk9sVVFqVEd1NkVCagoyWUZEV0kvdFBpNThkMi9ZMlRPWm5hN1RUOEc2VzJUZFlOTTlLVHdzSWk1a085TllmVm1IT2J3bW" +
	"t0akRFOEtLCkdPdnpzTFdSTjBFVUtOMlFWUktNYUNmV2hyZW4rSnRrbUgybTl5MzUyNFVld1ZKbUFyellaZnBNc1I3clVYK0QKczVDWDJ" +
	"vM1MyVU5WMEUzT2psNUF6OUhHYmdkc1ZwZXU5ajZwbTBhbVJlY0NnZ0VCQU1WcExwdDc0UjRhQUNuUQpzVTdVeVg3ajZrb1hTUzg0Zk83" +
	"eXJ1SWhPeE5MRWhpM3l6VTlCRzlPMk5CZEdlZHYxaDZYeE5mZWVDMkdCaGtJCnRvZkh3aGlYc3ovcEgrMnVzUDJ2QTRUQXJXNjRTNXJKN" +
	"zdDUjUyb3NtQUxkUUtKRG42ZjJnSEtNem56UFMxVjkKdEdyb1pMZ3lRaktDMUQvTnZ2SnhaR25wSTFtS3dGYi9panRKdGZqZkF1VjdxSm" +
	"lXZGdvcVBVSUUrSmMwWUVqQgovU3RMVnQ1MGlweW0zYnZwcSs1eTgzemtoNnlEcjRTVFJBRy9tdnJsQVEvVzVidEJ2ei93OUxGVXVGNlF" +
	"oUzViCktDU3ZlM21PWUJ5dWVDTE5tQ3h0YjUrQnRwUTRBbHZ6bHhudHFEQy8vbGp1SVF2RUVnTEoxMzhCZmlJSWRydUYKbUMyUmxkMENn" +
	"Z0VCQUtGRDlHSUVZdGtoZE1LUmo5S2xlbFZPU2JaVExxT0FLZG9SQXNlc05YZVI0dkpjZXR0Kwp5SzY0MG13cmkvVFVnSTVMYU1yU280V" +
	"1QxWUR0M3B3YVpDMVdxdlpJbngwWlhldFdaWXhNYUFxMG9WeS8rTUtSCjZtNXRDNC9zbS9wVCszYzZZTjF4UWQ5d1NUSHJwMndaUFBDM0" +
	"g2MGkzQlliZUtUY0JaYWgwMWpoSm5RSUFwR3IKMHZjbE1MbWc1RWJ0ekE0QWtqMENtaERab1E4TnhXbFhUYkJTWW4yLy9JbEFGaW1GY2Z" +
	"adW02YUJXU3ZZbEoyegpUUzRDYmVkSFlDang4TExyQmxPV3dPRE9HYzB2bS9JZStMdTJCQkdyOWlMWDFwTDZFV0tGTXRIUXEvTnMxTExw" +
	"CmhDRllNVHI1cldRUW8raTlYNXVKUGEyUVVnTVYvVmNzSGdzQ2dnRUJBSVFFVE1VNk1vYjBOTUx4czhtUFJOVk4KbDhMeEc0TFQ1VkNqR" +
	"mFHVDRia0RWQUoxMWFYMmhCNHdjL0xCaVZEWmJOcGQ5NXhaNzY3VE5VVG01c051dmdVQQpCMngzc21wUE1oampHLzMzUjg2d1lFYVY5c0" +
	"xzMEVNazRrOEM0SDhvMHdTNkJHTEZBZkRoVEVHZnpTYnBKdk84Ck5rMDRiYmFOM2t1TnJ2dHdObGozcG5jTDNQblE4UytJVkxlRVJTN3B" +
	"RS2FhdndrN3dpR1VkQXIyeFlsWEhzaW4KdHZ1Zkl3V2V5TFNXMElqaXZiUWZmZzJQeFg5QkkwMU1KNDhXdTlBL2x0YkFOdkRrZzFrcUZL" +
	"eXpUN041V0NhdApNbVJOUytuQ0xTeEN0eGM3cnBKZ01BQTBSaEZRZ25hcmF4Ukc5Nms2NG9UcmVreGd5VURFM09VMVg5UG55SlE9Ci0tL" +
	"S0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

func TestNewRedisTokenProvider(t *testing.T) {
	// ctx := context.Background()
	// defer mockery.Start(ctx, t)()

	data := map[string]interface{}{
		"str": "ABCD",
		"int": 100,
	}

	p, err := sec.ParseRSAPrivateKeyFromBase64PEM(privateKeyTest)
	require.NoError(t, err)
	s := NewJWTTokenProvider(p)
	tok, err := s.Store(data, time.Hour)
	require.NoError(t, err)
	ret, err := s.Fetch(tok)
	require.NoError(t, err)
	// It's tricky. int translated as float64
	require.Equal(t, data["str"].(string), ret["str"].(string))
	require.Equal(t, data["int"], int(ret["int"].(float64)))

	ret, err = s.Fetch("INVALID_TOKEN")
	require.Error(t, err)
	require.Nil(t, ret)

	_, err = s.Store(data, -1)
	require.Error(t, err)

	tok, err = s.Store(data, 1)
	require.NoError(t, err)
	time.Sleep(time.Second)
	_, err = s.Fetch(tok)
	require.Error(t, err)

}
