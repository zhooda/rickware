import random

x = "aachingandanyaroundaskbeenblindbothbutcommitment'scrydesertdodon'tdowneachfeelingforfromfullgamegetgivegoinggonnagoodbyegottaguyheart'showhurtii'mifinsideitjustknowknownletlielonglovemakemenevernoofonoohotherplayrulesrunsayseeshysostrangerstellthethinkingthistotoounderstandupwannawewe'rewe'vewhatwhat'swouldn'tyouyou'reyouraachingandanyaroundaskbeenblindbothbutcommitment'scrydesertdodon'tdowneachfeelingforfromfullgamegetgivegoinggonnagoodbyegottaguyheart'showhurtii'mifinsideitjustknowknownletlielonglovemakemenevernoofonoohotherplayrulesrunsayseeshysostrangerstellthethinkingthistotoounderstandupwannawewe'rewe'vewhatwhat'swouldn'tyouyou'reyouraachingandanyaroundaskbeenblindbothbutcommitment'scrydesertdodon'tdowneachfeelingforfromfullgamegetgivegoinggonnagoodbyegottaguyheart'showhurtii'mifinsideitjustknowknownletlielonglovemakemenevernoofonoohotherplayrulesrunsayseeshysostrangerstellthethinkingthistotoounderstandupwannawewe'rewe'vewhatwhat'swouldn'tyouyou'reyouraachingandanyaroundaskbeenblindbothbutcommitment'scrydesertdodon'tdowneachfeelingforfromfullgame"
x = list(x)
arr = []

while len(arr) != 256:
    lv = random.choice(x)
    rv = random.choice(x)
    if (lv + rv) not in arr:
        arr.append(lv + rv)
        
print(' '.join(arr))