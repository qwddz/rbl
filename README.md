## RBL Checker

The RBL package checks the IP Address in the Real-time Blackhole List - a list where malicious addresses of email clients are entered.

## Build
To successfully check the rbl lists, you will need to build a binary with cgo disabled. This can be done by setting CGO_ENABLED=0:

`CGO_ENABLED=0 go build -v example.go`

## Examples
With *default* rbl servers:
```go
package main

import (
	"github.com/qwddz/rbl"
	"log"
)

func main() {
	res, err := rbl.New().CheckIPAddress("91.240.141.42")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Status, res.Errors)
}
```

With *custom* rbl servers:
```go
package main

import (
	"github.com/qwddz/rbl"
	"log"
)

func main() {
	res, err := rbl.New().CheckIPWithRBLServers("91.240.141.42", []string{"dnsbl.spfbl.net"})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Status, res.Errors)
}
```

## Default RBL servers
* access.redhawk.org
* b.barracudacentral.org
* bl.spamcop.net
* blackholes.five-ten-sg.com
* blacklist.sci.kun.nl
* block.dnsbl.sorbs.net
* bogons.cymru.com
* cbl.abuseat.org
* dev.null.dk
* dialup.blacklist.jippg.org
* dialups.mail-abuse.org
* dialups.visi.com
* dnsbl.antispam.or.id
* dnsbl.kempt.net
* dnsbl.sorbs.net
* dnsbl-1.uceprotect.net
* dnsbl-2.uceprotect.net
* duinv.aupads.org
* dul.dnsbl.sorbs.net
* escalations.dnsbl.sorbs.net
* hil.habeas.com
* http.dnsbl.sorbs.net
* intruders.docs.uu.se
* korea.services.net
* mail-abuse.blacklist.jippg.org
* misc.dnsbl.sorbs.net
* msgid.bl.gweep.ca
* new.dnsbl.sorbs.net
* no-more-funn.moensted.dk
* old.dnsbl.sorbs.net
* pbl.spamhaus.org
* proxy.bl.gweep.ca
* psbl.surriel.com
* pss.spambusters.org.ar
* rbl.schulte.org
* rbl.snark.net
* recent.dnsbl.sorbs.net
* relays.bl.gweep.ca
* relays.bl.kundenserver.de
* relays.mail-abuse.org
* relays.nether.net
* rsbl.aupads.org
* sbl.spamhaus.org
* smtp.dnsbl.sorbs.net
* socks.dnsbl.sorbs.net
* spam.dnsbl.sorbs.net
* spam.olsentech.net
* spamguard.leadmon.net
* spamsources.fabel.dk
* web.dnsbl.sorbs.net
* whois.rfc-ignorant.org
* xbl.spamhaus.org
* zen.spamhaus.org
* zombie.dnsbl.sorbs.net
* bl.tiopan.com
* dnsbl.abuse.ch
* ubl.unsubscore.com
* dnsbl.dronebl.org
* 0spam.fusionzero.com
* 0spam-killlist.fusionzero.com
* combined.abuse.ch
* drone.abuse.ch
* spam.abuse.ch
* httpbl.abuse.ch
* rbl.altaria.com
* orvedb.aupads.org
* fresh.dict.rbl.arix.com
* stale.dict.rbl.arix.com
* fresh.sa_slip.rbl.arix.com
* stale.sa_slip.arix.com
* badhost.stopspam.org
* list.bbfh.org
* block.stopspam.org
* list.blogspambl.com
* bsb.empty.us
* bsb.spamlookup.net
* rbl.choon.net
* query.senderbase.org
* v4.fullbogons.cymru.com
* tor.dan.me.uk
* torexit.dan.me.uk
* bl.spam.deadbeef.com
* dnsbl.stopspam.org
* rbl.dns-servicios.com
* dul.pacifier.net
* rbl.efnet.org
* rbl.efnetrbl.org
* tor.efnet.org
* fnrbl.fast.net
* 88.blacklist.zap
* random.bl.gweep.ca
* hul.habeas.com
* forbidden.icm.edu.pl
* spamrbl.imp.ch
* wormrbl.imp.ch
* dnsbl.inps.de
* rbl.interserver.net
* any.dnsl.ipquery.org
* backscat.dnsl.ipquery.org
* netblock.dnsl.ipquery.org
* relay.dnsl.ipquery.org
* single.dnsl.ipquery.org
* spamlist.or.kr
* admin.bl.kundenserver.de
* schizo-bl.kundenserver.de
* spamblock.kundenserver.de
* worms-bl.kundenserver.de
* dnsbl.madavi.de
* c10.rbl.hk
* bl.mailspike.net
* z.mailspike.net
* dnsbl.forefront.microsoft.com
* combined.rbl.msrbl.net
* images.rbl.msrbl.net
* phishing.rbl.msrbl.net
* spam.rbl.msrbl.net
* virus.rbl.msrbl.net
* unsure.nether.net
* ix.dnsbl.manitu.net
* netblock.pedantic.org
* spam.pedantic.org
* rbl.polarcomm.net
* all.rbl.jp
* dyndns.rbl.jp
* short.rbl.jp
* virus.rbl.jp
* rbl.abuse.ro
* abuse.rfc-ignorant.org
* bogusmx.rfc-ignorant.org
* dsn.rfc-ignorant.org
* postmaster.rfc-ignorant.org
* dynip.rothen.com
* dnsbl.rv-soft.info
* dnsbl.rymsho.ru
* all.s5h.net
* exitnodes.tor.dnsbl.sectoor.de
* query.senderbase.org
* bl.score.senderscore.com
* problems.dnsbl.sorbs.net
* proxies.dnsbl.sorbs.net
* relays.dnsbl.sorbs.net
* safe.dnsbl.sorbs.net
* nomail.rhsbl.sorbs.net
* badconf.rhsbl.sorbs.net
* l1.spews.dnsbl.sorbs.net
* l2.spews.dnsbl.sorbs.net
* recent.spam.dnsbl.sorbs.net
* new.spam.dnsbl.sorbs.net
* old.spam.dnsbl.sorbs.net
* geobl.spameatingmonkey.net
* backscatter.spameatingmonkey.net
* bl.spameatingmonkey.net
* fresh.spameatingmonkey.net
* netbl.spameatingmonkey.net
* uribl.spameatingmonkey.net
* urired.spameatingmonkey.net
* feb.spamlab.com
* rbl.spamlab.com
* dyna.spamrats.com
* noptr.spamrats.com
* multi.surbl.org
* dnsrbl.swinog.ch
* st.technovision.dk
* rbl3.tfoms.perm.ru
* opm.tornevall.org
* residential.block.transip.nl
* rbl-plus.mail-abuse.org
* rbl2.triumf.ca
* dnsbl-0.uceprotect.net
* ubl.lashback.com
* virbl.dnsbl.bit.nl
* virbl.bit.nl
* dnsbl.webequipped.com
* db.wpbl.info
* blacklist.woody.ch
* uri.blacklist.woody.ch
* bl.blocklist.de
* dnsbl.zapbl.net
* rhsbl.zapbl.net
* zebl.zoneedit.com
* ban.zebl.zoneedit.com
* bl.suomispam.net
* spamrl.com
* truncate.gbudb.net
* dnsbl.spfbl.net