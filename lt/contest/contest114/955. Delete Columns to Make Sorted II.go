package main

import "fmt"

/*
We are given an array A of N lowercase letter strings, all of the same length.

Now, we may choose any set of deletion indices, and for each string, we delete all the characters in those indices.

For example, if we have an array A = ["abcdef","uvwxyz"] and deletion indices {0, 2, 3}, then the final array after deletions is ["bef","vyz"].

Suppose we chose a set of deletion indices D such that after deletions, the final array has its elements in lexicographic order (A[0] <= A[1] <= A[2] ... <= A[A.length - 1]).

Return the minimum possible value of D.length.



Example 1:

Input: ["ca","bb","ac"]
Output: 1
Explanation:
After deleting the first column, A = ["a", "b", "c"].
Now A is in lexicographic order (ie. A[0] <= A[1] <= A[2]).
We require at least 1 deletion since initially A was not in lexicographic order, so the answer is 1.
Example 2:

Input: ["xc","yb","za"]
Output: 0
Explanation:
A is already in lexicographic order, so we don't need to delete anything.
Note that the rows of A are not necessarily in lexicographic order:
ie. it is NOT necessarily true that (A[0][0] <= A[0][1] <= ...)
Example 3:

Input: ["zyx","wvu","tsr"]
Output: 3
Explanation:
We have to delete every column.


Note:

1 <= A.length <= 100
1 <= A[i].length <= 100
*/

func main() {
    /*
    fmt.Println(minDeletionSize([]string{"ca","bb","ac"}))
    fmt.Println(minDeletionSize([]string{"xc","yb","za"}))
    fmt.Println(minDeletionSize([]string{"zyx","wvu","tsr"}))
    */
    //fmt.Println(minDeletionSize([]string{"xga","xfb","yfa"})) // 1
    //fmt.Println(minDeletionSize([]string{"bpubea","dwgjgd","dqbsih","tzsegm"})) // 2
    fmt.Println(minDeletionSize([]string{
        "vvxejcrulnxamfxbkqlnpiahmbuvltqjqtwocaxapyicraeufsoppnvgjzviemxjurvmqgudlzhgupsvsxzzfgfkkrdalgnqommw",
        "lqakiuozyiuuyjfbjsfkbttmrnctgexxsizfraxfspvffrfsangdsbueyqfpkpqieqaafhbaocbzajmnidtnwqgkprixqzdfblwp",
        "majviidklyhfrixjwdhkbwoyuoldrttduhfokariaeqwgyzscjoibzbhsjtfcefzlwnskmzltizvbigduavvkvfjohjbcslgvwcy",
        "ddichavlcenxkpnbqkiewybdkmsafrnsbtnovatldrxlqcvgmeawusyflznkjbdablvvlbotyzydzbpouwbxotvbrpssoafyaran",
        "mevbayyazaqbxlnogwzzzsfrziuyrzinangdearuztvhscqdotorjdaoempuzrydtmgylzkcfcdhtuovsueuyfmjdyyqaevempxv",
        "swxumgueqqyztrcstcpsziirmqeugrndptzdfaovvvcomhljsvpoghfhcxnzetmruxrhmvvpljewvojhuztzjesommrnwhvgwrcf",
        "oouxsczlswqduchqbexhcammderstuxjludxgalvdyvozzjlrkdvciahtmoqiofieftsaioawvswqqvijnamrbbanxpgffixzvkv",
        "bcxmaudsiwrfcetjotmjbxrzxeuterckvcdcqbncdpyybqpuibptfwaaeyglghhpynyagjayiryyzugkscxlhctekocopsqfibai",
        "dmtjaomvcvtmcmyprrpkthotiouicaiigrmarcbgwpgkaxfgogskgilastwjyrkwlbqtntnxgkvfpfdypkxbrseiriqtnwqteooa",
        "xrkvkjseixbgjfbbeswzygmncfriynnjbfthucsjtbkhsiehqmchcfqzjpqajnfdxbpeaawtjukuaadunhdxtwjeyzlibxanlgtt",
        "dcbzjujzdtfrzvmectlxxnzrbrjeoajfqnmxhcvltpcltpawljpmjtdwtfarmgwbhybzljmboopukijtyxjpptgfkexhohrmpldf",
        "xeasfysgofkerpfvhufzmcssnprkpwfnbqxypcmsroeeesfcocbrkacqkuimbfoooiyczxlkmptfiocbwbjpkgihawrufyovgxsu",
        "tnnvfwbuavaifeieajrnvgwpiphxwwapweonmcdyxccirwdlqpngvckdduhuosomoekppxoksilhvmpdhpmajtpxvipeenmkkafd",
        "kjsnbzgyzozlfsxswyllbesjggiiicicnjadecgzgwddmfdrlyayuhmmuuugcnkbotxlrcedjnmlrkikxsjqqgejqkdglkazvjic",
        "syxvebchdqgxrqfdgepzkwalatgdmopigdvyhdkgkoenrjyjojrnouysjfzbzhluavvvzgigcyprbttdbdmvwunjzcmokicbmwig",
        "jobdhujosadevfjpvjblhkrtwrcgigbhadteadbkuubveqjqktmnetrqqztjkajjzgwfxtuzttdxqffwnltynwxmthzndzweyack",
        "sjlvidfovhreachpjlohhfiwhckivfnzkycmzdbljbqsgieohgcmfpugyoyiuvrrnmlbwedkkiomtahiunkxbydwhgnrpaqocvha",
        "yuuobhxdqkefwsztcrsutpqrrqdtgcivycorydzsirphcoltyffxxiwvykyvhgheuabmujfgystnxecahzcxeeqleuqnpejtlrtb",
        "odoqphvufxglfojqhdhagryhlrdlphcrvrlxmeenusjojhayyuryjbqjimpbrskniqpozpwxvnxzzawjxaidfunanyfnpkuqvgcg",
        "tnedwjomyuwgyuonaswjlvwnoytgbmawavdjaewzsgxgrvqlzyojbxtloxxqggpleawbetmnqmtqwvfvittukpegmccdvbfijvwc",
        "tmtmjvrbdkerpukrcumpctmmxzoejxmpuepkuftnmieoieesueulddkchrrnovtxoqtubxwqossmwyvgsrgqansjdqcbvftahhgn",
        "rttugtufedysqbxcshtnpsrcdgiqiqqoqdlhxgkdektndvvrrpizdysvxwcsmsvrhawowloiumbacceurroxbzugcqqwdwseiqcg","nlrbahdkkzhlbjxaevpdfbxpfzklvheggwxsmgnjgnopqrmgdvhemsgaffrfyogcyfqqtzyuasmlfzhxlrpuvhypoggytwzbkdak","pfsqgoawrzmtppneapcmacwsrfxfxnxgbsznygloeshnfftsnnirispfgsnnmdbwknjkslqrbzswjowsufnunazstbgggmjjaaal","dbzwiiikwpxdamatqayeqtovwqbncnhuvujcdhcahpazjbuavofknubgkkezaiuvkyxbqihpszfvidrpolmtzfpdbubngeufukze","hiutwttchtsxstfmonvrmfswlqhtkcelwrkxlhmkotjjazdlryuxbnerlxbegtsfyweilekuqqnuyzkljtjsajisdlficxpyodac","oidzdupsecdfawqbyofdwglamhtveotbrwguzhhtafdxqugpibjlgfcykeinoyeddeuzbsuaovzvcyagbgqwpnqjurrmgucxxvzr","ttfequhoomcbckhriplkrjldjjgbuosozsxhdidapwgmiztbjriycuzhtwbablrvegmntvihpxwvtohcqmrpbpdeofexcffyisbs","rxdrwsuixnvxwxxvxqdamenqsmjqakulfajsxiicpaifpqykqfiqtqumjkmegszbxmrjfajxohkvbxntzjnkgqutuhgwkazdqcdo","qtmawynhcbfvmkmfejmrwfforuetbotadjxhhiahwysegxklkssgmgpjyjzivlxzhybjvpoxdzcyqsdtjfcniqwrjbeveamhfrun","fbujppkgzjxpkzezgsmrvbrwtarwyhtwbiyoziklmefyptiapavsnwcfilgcnobxtykhdqbabagvloxmjwpoqfhxkxtruklzliaj","lvsxjycvwztduyyiiemkfnbezdirrnpwellteitnwaoctlihcxsxbnydicytcltfewazeyvbbsyztsdmyjkfmbeanovqtgojmpfi","lhlpqymecsqtctvidukotjbdlpfxfvoedfgmrirqndvaehyxexnbypoeugbuuwtvvbziulwbqlbylljehaxmhvwyvbatccmglhmv","jjzradoxekdczywpwepqpcafjoiffjtulxqlsikulgmevkjlypwbedxrfmjltlrqzsrzxlwmtfculzsomxrehgbkrmamrisxpwfp","yspowyvusaurikqpoixlzqmiabhgkvszveuxrjxcondwnihowwtffrmcthdnksepcspmteignoxhmetkddmndtmgytqjdmepfhfi","csctrlhcvsshxiipwmduhnrkiupqszquwdwxbjnouzreyjyqeljbsekdunzmhfgnjvvhojevgoynwpxvlsilcrbtisdynbjwdmov","jlhaprojrzbhhtotkhdfqkspiubdnzomurdbrkkdwrstpexaligmjfmjyftuwhkqeludlutrrndsmsbylbmpdvikmombgpdpgwij","bjbvltxijnsihgjurnwfgywgrebqnxuitzqvdkzootvmjhwzotqfyxepypfdvykzwidscykpmafucybkqnkvmgwqjrfzzixagdoe","xtfcxgazxutdbazevijusdnkabhnkxypibfujldakzuhgzgneqwpumowqhirpqvmnmvxeeuixuoqrywjatlagqmccpylzwcogppd","gykduitgbuxizbmdluoqzjspgrudcihhdfkvolpdbaodqxmdradsvbxzpqobwijvbtdaauvhxmyzogzesmxjjtuyclznvrtjvvjo","nkbpmqmghhqlrhwnfgopfiubugrumefwlqqrzlwiseclgiqzkoqfljqfjqfqsvlmxdkhdfndtgvmuxpbfwdzkasverryenzfqxwi","vwdhcgxazszewfahphcgaruzjmswzvrrtuqoxluibgmaziifyceelkcpgemlciewmvewcpgoqxueiyrysjrpiepkqcmnbkehdorn","gnqqyxsnoacdqafifihlvgymormytbqvcklydlcrcwbeuinvecndshiaothpobikbdhxpujyqxaipkkyrxfanwxykpgayexfexzf","pqbhqleuzhznwedawihmkqsibsshslbmrecvpmxbynqlbhcsaclbixbphfndiaepuqchgfxpoxvvfopjmiwevfsjtjynlpjbchve","uyxrklfqbbbpdclgblgkevrqgmmvwpuxqeqhymkcnvikfpcukuznocokolqegtchtzcboebircudlmlguyceqdtfeveovxoppsex","jfqhlqkrlgyitnwzcfofkmmzqduraaxydaafdmeorhmalzxdaakrlfemruoxqyaaluwodznukscrnnsxlnhocgollnfbunnkmcvs","fimpptwphjomdoehcpovclnoshpufjscxcjbsmcusxjilgikoqqqjrmynxmrdlijvkwdkolpieurtyvxomdkzfsrgtrgjjxcctup","nldvamkchsjbpqecawtqocstjoiuhiesvxkeomovazwheycwrrhaofsesgkiiislzefleabsqfpspdwkdxvddrrffqrcmcxijtkb","euyxnjjezrloluwwngnpstoqubtlatkcotkvhncajvaeinpqnzozglctrvwvkbpygsjbtnyyyrbnkjlqkxhciuzglilyxnihjhjm","aiesxbgplcbtumyxurccirxsufuleardkvjppnsljfknqqiordouqxysneqcpcymppjrsidyjwzizxljjxpbmuhtrwbxsclfdqwm","rmtfualnxcjnsrsrwtxknyzoqxejeqcglwzmjnbsgrlefrgrletnucutlmfcteyvzbiglpmzviyhyqazvhulxtqqlnqlfwmcuxxy","awrxgheeaoszkfddwadtohjqfayvgafqmcovlnpshydzfecfhazmhqhzyjldwljpiesdtfywijnninczrctinrvnbckazagjaaqe","oiwkfjyozyhsirieidacsdzdaojdikpfdydjbnaxvelruqaouhifvvqfcjjkdmrrfrpdtxyapannqicvkwvvgalzqxgouoajfpzn","lsmzdlwfbvwathgnscdgaltlqafumdukvarclogejgozolampqewjglziaveeiszgemfssiquhcuqkhjelrwkygcccbzdzyqtloo","fdtxqknyrsanwjwdkcsojtsbjvkmacrsqvsirojghqgegnvliyxhqdmescoyvzvahdfhogjcyppkdwekhdhsjfejaocclmfjogzw","xvbutyybaqhvetstybermemdanmsjsnfkhibkoauyphnwggncyrlicryfhsxehxrcqlnbfcyvyjopotalqevayptkzxpwunwvrba","rmjjkhpgdlmfepsbvsdfwuoeiajqoelxkleapogwxmnyrqefblxiznwchridsphoyudlypcxyakmzgottlpartkgkrmmwnmzvglz","egvamcvftjfdhabkkyldlbvvupwxflgmcjipgpdgfthxatoqgfvfizsiacjiymcwmsefguxhkpmpyijxekjbnfuonhkuolvfabus","jcysacpeljlcnnmlaxnzjtadwphppngexewlcpnowbltcdwwxzgpqnkogfgfslqfycetghgchveecjrmrxpjghovhsgxckkxzkyc","akyndnocnvrbizjplcswjyyapnpmzwuhmubqnpkzfldhzazkmkbiwlnokjsispxyzclnxkxsvfegygnqnfpqjbjeedxrlejefgqu","ljpnsefnaxgntothoyoexdipzmujbaispfdjaqcbqgjsedyadgjlvokyyvnmjyvdklkrlixsdvhfedbtussgottmgmmgnfgsuzkr","dfahrhmtrnmdsckcswccctmjbaxsgorcqvqjsqaguesljphonqoyjwfpbxlfcvcxptuumxdfhyrvlxjlewvwnmbvsdreswvdazeg","pyhemaflwmgwdpntkiaxoankebqpsaajelyrkqxlrfnvwdaecstieumnhisyavdmjhykxdswlclikmcxjpycnttepqlxgmapfsyk","bskiaplfbfopgeledgbtnhjugwnqkubdftswcqrpnpaojwipwghfmaxxvgqrbgkiamhklmpqlshbuwrahiuicimgtquezvtdieju","dajfjqqeqonwyvuyostjlyksuqmgdqorptvqoqquijccopxsewwwxxktwuaiqpbenszqzmvrxvbqkmxidbqnwoxjpzelbnqwhjuh","mjaqekpjoetibplniwilxiiexrljgicgmkkueqouptkvbhlzdlfbtcbdsfmzqzanszwktvoniazplvkxbdfpxsiubrgxpobdsfta","zfkqgakxkjhcirvvhwrhinaybrymxdqtleqqlrhayaxqehgqihbcdidwozrgdpaezpzcvmbtgpennwrftagsbxcnrtzvriauokmi","jdbxxhdirtrrruotajmvzaiwpvsbndrnywzbsrzbkkdvybketfvaltfazlvppsbbcbzjzskictpshyrztnnqnlvcxlmauujbpawc","zkbbxpufdgijsziubgysvtghzgfjzylfwhnigrrertqptsoormkhajtlmuzlbukgwsnigurlzovmtihyfroiasrpqsvbiswashge","nlxvrmpgsdbkplnprypkrvfvuvwlnmpmtqsrirlpahzjxawbahrkehbcbeorefhdhfmndtcshceoovmchkjuzygqyqbpvibiueju","tqkckzvcawhgbwafeueqdokidutnzoblajwqlstmfaoumjfaqsqzfbzgwnwdrwjfvgulzwsughjpazwgjofslnrajoetxxpcdhvm","nfkquuxgsqgtpaoqqmvzhkvehsdeoghgukmpdszszwneohxzelgnmqfqlkwxerlxnojiblbnigfawxwxidlzzhomcrbvvdhhgisc","gqdeibxktuaqdheoegcidmqdkvddwhckwdjoxtsgrdnfvboqdxalekkqbvwjnfwswjmddnppqgarnhaofrweinsfybfcmgelrvkq","jljgahkwfqwofzgmvhjarqrqpordhumrawqchtshdocsrzjlugnlflzhdqkizvhcheppgxzujiockxghvbhmpwqgqkhglcfjmyuh","ajoyvegffjencywgixdmbkvfflvhjyvwynrwmtzkzkdncppzicvahvueygwnzhovbfxvqehxufbppwgrlsivokozbouhyojwbhwb","iekpdgmeaihgkvvxobkwoapvomuobqytcywkqtcmpphaevnhvgzqxillcanxzzuszfcemdzhdrnktkfljduzxoviqpucqeejyzor","cbhonhiqdjyisbnyvixhmgbhmwpakzyeqsfwcthpnecitkmzhtjqmfzimlkyqxraxotmrolyeciwysmyopzxrdlifzpbfyvgpmps","bqqnpzynrwlinhydagfkpmblccjlqymmpzytetrqbrwwkcmofbytcvephmplyjuoyqluflqoexgbrstksxcezkrkfofppuyyjmrd","xinequgxdwrjcanssohvbakrkaoahgdnfuprhumdyjtigkfafbhzohpgiwbiggoueqribgwnmfzsdjygdaqgpdnghejvqyerhtcn","daxnbubhbvicqiotyscvprrmwrcirkatsymwsuddszdmcwobdtonvppphmyhisduwicerndyrmviugmjmsecqfjfuilitbprekpb","fxkpfnwvshjhbsjrdpaarwvoyxubpzatxevbxubgdteqaovhphsohvnrwhjgrfyrlgdciqkskvalxfzeqkiokvzqfiqwzeygqklz","youflzgvieowuzcljmqxwstpibfqkynsxljscuyltpppcvfbslcqocxuyztickgnyexgqyravovkykpscvzyhfdwjtdyyleejfav","wkmzqnkvlxrzjnjlvnnqlpydvvyavzhaylfjjurpogexwfipzetvewqagsltqugadpsfalwyhhcxstyqpeddrszjyuqvkjmrmcfr","pnmreilozjohbhgentgtrsowcmoxcfwpxazygvqeocxiapqsxzyctlebqmhwqxpnpnorksiqmlwbgzuaxxmtyelmxihnmtwqexvh","mgrjsweqfamcsrcxkplgyxivnhleovcqdznyfvnfqhtpxdbutfocffgdoplvntuixroinocjkhurwywnoocqceseettdaidkfpgk","kstvfondsurxhxhtuiooyrbzknrkcpcxzwdzrveikjyyaazidyuaniohwejiacnekknhkbwhkrmtwwowrgbwmncaikafkadizmtg","nhgxqrasosotxncapimdjiyllndiaflwqublvvnkqkjdmckidcygyhwbjwmoutzhyllumsapjenpmkqywtitfywsleeaqcfcyazh","oyzmysmwmkvphyoxzvqtjmaykkmdmtqxlfpjgvklsnvanyxvhoxzpuuxyeeubozejyxbibxdfcfzvqmtxmyxgxquaouiucehxvgn","hakggvicoivhvgxoesqwjkrepbfplkwfurturwbbejtpozgxttnnkipguegbesvdhglhpfhibqhoqtxgwdgqjdfirmcukvdgswjy","kfdhfecozbxxwynqntdwapxhsxnkzjcuotriuwwdgnoinqeohkmxctrcygqppdjuotwvayefbglmvggbqhsyonasdkibofxaekhk","zdanrjtgzjqrmmdvpgiaayjzmonbtuqdojiwcwrfncciemecubruaxxfwuqqvuflprqqyybhllelghixpkuubmzwcpdyjglryeoz","nyfwepwfcmgtkfmuwnkrsdwwwrxvgjhwkihiixrbagjjkkysyiiubtapzpascckcngqygqkaovzkjleqhguxwgexskydpupehhln","yousfqsojqtyeoirmrjzfehlehpmoxyprkiduxsgxcgaddoqxstyarcavfappojjpsxxivcrdzpofitzunfjdmidumnnkcnxczcs","wazcovqsnqizkwvbufknrzfiyaufckaifuclgxeljhhjfizfieopshkwmnrceegdamnfewooecnomggcsmciynypkkcqxezjyjbl","yrryieqyeefiubgywiugmhlwcecyisevwawsexystbsikfjbwzgljdpldorogwyvcoqreqhloaqiogbanukrhmccbubuvzgvafqr","gtnbojfqcnyhcodzzvrerqcwaogevcczrstelxssrqabvqnqglwuqgjwqgjxjvojisrhtrjogumpojtlnxbgznjftwacshuutzyq","oovyqqzzxreckjykwwkhippffulmjjowauswcxutmujujibyfrxwabfjbyhjkxzlotaecsyziyytxebtinmwcwxgdvbzrbiquwag","jzigqkybsssewtjmmeheiwccfweiufeqgeqrnxvyjzxhacxjccxefyjkuftkiaablvtdwivtfpkovprbripodntdlclxdntwkayv","jqskcwpdcpunmncsynowjifmhdpzosehgbbdbypveqqczdptqbagrfhunclqpehwhdaujkamecpykuabvajpevysrpkiztasdtsr","hgludaxhhjdagglsjudkefuoilctomlfldxxwzptbsmqcealippibzydfokzddfaxbcczgwxkkbrzbgqptwfgagvlniuqqagrcen"})) // 1

}
func minDeletionSize(A []string) int {
    if len(A)<=0 { return 0}
    row := len(A[0])
    prefix := make([]string,len(A))
    ans := 0
    /*
    一列一列的看，看当前列是否可以纳入进来
    每次都看当前前缀：全部看完所有列即可。
    */
    for i:=0;i<row;i++ {
        cur := make([]string,len(A))
        copy(cur,prefix)
        for j:=0;j<len(A);j++ {
            cur[j] += string(A[j][i])
        }
        if isSorted(cur) { // 如果当前前缀是有序的，则保留下来；
            copy(prefix,cur)
        } else {
            ans += 1
        }
    }
    return ans
}
func isSorted(A []string) bool {
    for i:=1;i<len(A);i++ {
        if A[i-1]>A[i]{
            return false
        }
    }
    return true
}