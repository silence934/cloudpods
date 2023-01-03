package models

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"testing"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/pkg/util/secrules"
)

func TestSecurityGroupRulesCompare(t *testing.T) {

	inDefaultRule := cloudprovider.SecurityRule{
		Name: "inDefaultRule",
	}
	inDefaultRule.Priority = 20
	inDefaultRule.Action = secrules.SecurityRuleAllow
	inDefaultRule.IPNet = &net.IPNet{
		IP:   net.IPv4(1, 2, 3, 4),
		Mask: net.IPv4Mask(255, 255, 255, 0),
	}
	inDefaultRule.Protocol = secrules.PROTO_TCP
	inDefaultRule.Direction = secrules.SecurityRuleIngress
	inDefaultRule.PortStart = 1
	inDefaultRule.PortEnd = 3306
	inDefaultRule.Description = "i am inDefaultRule"

	outDefaultRule := cloudprovider.SecurityRule{
		Name: "outDefaultRule",
	}
	outDefaultRule.Priority = 20
	outDefaultRule.Action = secrules.SecurityRuleAllow
	outDefaultRule.IPNet = &net.IPNet{
		IP:   net.IPv4(1, 2, 3, 4),
		Mask: net.IPv4Mask(255, 255, 255, 0),
	}
	outDefaultRule.Protocol = secrules.PROTO_TCP
	outDefaultRule.Direction = secrules.SecurityRuleIngress
	outDefaultRule.PortStart = 1
	outDefaultRule.PortEnd = 3306
	outDefaultRule.Description = "i am outDefaultRule"

	src := cloudprovider.SecRuleInfo{
		MinPriority:    1,
		MaxPriority:    100,
		InDefaultRule:  inDefaultRule,
		OutDefaultRule: outDefaultRule,
	}

	dest := cloudprovider.SecRuleInfo{
		MinPriority:    1,
		MaxPriority:    100,
		InDefaultRule:  inDefaultRule,
		OutDefaultRule: outDefaultRule,
	}

	rules := cloudprovider.SecurityRuleSet{}

	length := 100
	for i := 0; i < length; i++ {
		rule := cloudprovider.SecurityRule{
			Name: fmt.Sprintf("rule[%d]", i),
		}
		rule.Priority = rand.Intn(100)
		rule.Action = secrules.SecurityRuleAllow
		rule.IPNet = &net.IPNet{
			IP:   net.IPv4(byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255)), byte(rand.Intn(255))),
			Mask: net.IPv4Mask(255, 255, 255, 0),
		}
		if rand.Intn(10) > 5 {
			rule.Protocol = secrules.PROTO_TCP
		} else {
			rule.Protocol = secrules.PROTO_UDP
		}

		if rand.Intn(10) > 5 {
			rule.Direction = secrules.SecurityRuleIngress
		} else {
			rule.Direction = secrules.SecurityRuleEgress
		}

		rule.PortEnd = int(math.Max(float64(rand.Intn(65535)), 100))
		rule.PortStart = rand.Intn(rule.PortEnd)
		rule.Description = fmt.Sprintf("i am rule[%d]", i)

		rules = append(rules, rule)
	}

	src.Rules = rules

	dest.Rules = make(cloudprovider.SecurityRuleSet, length)
	copy(dest.Rules, src.Rules)

	common, inAdds, outAdds, inDels, outDels := cloudprovider.CompareRules(src, dest, false)

	if common.Len() == length && len(inAdds) == 0 && len(outAdds) == 0 && len(inDels) == 0 && len(outDels) == 0 {
		return
	}

	fmt.Printf("common:%d\n", common.Len())
	fmt.Printf("inAdds:%d\n", inAdds.Len())
	fmt.Printf("outAdds:%d\n", outAdds.Len())
	fmt.Printf("inDels:%d\n", inDels.Len())
	fmt.Printf("outDels:%d\n", outDels.Len())
	t.Fatalf("failed")

}
