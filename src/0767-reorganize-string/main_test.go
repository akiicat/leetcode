package main
import "testing"
import . "main/pkg/testing_helper"

func TestReorganizeString(_t *testing.T) {
  i, o := "aab", "aba"
  T(_t, S(i), S(Valid(reorganizeString(i), o)), S(o))

  i, o = "aaab", ""
  T(_t, S(i), S(Valid(reorganizeString(i), o)), S(o))

  i, o = "zifrfbctby", "bibrftfycz"
  T(_t, S(i), S(Valid(reorganizeString(i), o)), S(o))

  i, o = "zqugrfbsznyiwbokwkpvpmeyvaosdkedbgjogzdpwawwl", "wzwzwzwawabebebsdsdvdvgygygfkikjklomonoqprpup"
  T(_t, S(i), S(Valid(reorganizeString(i), o)), S(o))

  i, o = "tndsewnllhrtwsvxenkscbivijfqnysamckzoyfnapuotmdexzkkrpmppttficzerdndssuveompqkemtbwbodrhwsfpbmkafpwyedpcowruntvymxtyyejqtajkcjakghtdwmuygecjncxzcxezgecrxonnszmqmecgvqqkdagvaaucewelchsmebikscciegzoiamovdojrmmwgbxeygibxxltemfgpogjkhobmhwquizuwvhfaiavsxhiknysdghcawcrphaykyashchyomklvghkyabxatmrkmrfsppfhgrwywtlxebgzmevefcqquvhvgounldxkdzndwybxhtycmlybhaaqvodntsvfhwcuhvuccwcsxelafyzushjhfyklvghpfvknprfouevsxmcuhiiiewcluehpmzrjzffnrptwbuhnyahrbzqvirvmffbxvrmynfcnupnukayjghpusewdwrbkhvjnveuiionefmnfxao", "eweweweweweweweweweweweweweueueueueueueueueueueueueueueuhuhuhuhuhuhshshshshshshshshshshshshshshshshshshshphphphpcpcpcpcpcpcpcpcpcpcpcpcpcpcpcrcrcrcrcrcrcrcrcrcrcrcrmrmrmrmrmrmrmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmgmgvgvgvgvgvgvgvgvgvgvgvgvgvgvgvgvovovovovovovovovonononononononononbnbnbnbnbnbnbnbnbnbnbnbnbnbnbabaiaiaiaiaiaiaiaiaiaiaiaiaiaiaiaiatatatatatftftftftftftftftftftftfdfdfdfdfdfdfdfdfdfdfdydydydydyzyzyzyzyzyzyzyzyzyzyzyzyzyzyjyjyjyjkjkjkjkjkjkjkjklklklklklklklklklklklkqkqkqwqwqwqwqwqwqwqwq"
  T(_t, S(i), S(Valid(reorganizeString(i), o)), S(o))
}

func TestReorganizeStringFill(_t *testing.T) {
  i, o := "aab", "aba"
  T(_t, S(i), S(Valid(reorganizeStringFill(i), o)), S(o))

  i, o = "aaab", ""
  T(_t, S(i), S(Valid(reorganizeStringFill(i), o)), S(o))

  i, o = "zifrfbctby", "bibrftfycz"
  T(_t, S(i), S(Valid(reorganizeStringFill(i), o)), S(o))

  i, o = "zqugrfbsznyiwbokwkpvpmeyvaosdkedbgjogzdpwawwl", "wzwzwzwawabebebsdsdvdvgygygfkikjklomonoqprpup"
  T(_t, S(i), S(Valid(reorganizeStringFill(i), o)), S(o))

  i, o = "tndsewnllhrtwsvxenkscbivijfqnysamckzoyfnapuotmdexzkkrpmppttficzerdndssuveompqkemtbwbodrhwsfpbmkafpwyedpcowruntvymxtyyejqtajkcjakghtdwmuygecjncxzcxezgecrxonnszmqmecgvqqkdagvaaucewelchsmebikscciegzoiamovdojrmmwgbxeygibxxltemfgpogjkhobmhwquizuwvhfaiavsxhiknysdghcawcrphaykyashchyomklvghkyabxatmrkmrfsppfhgrwywtlxebgzmevefcqquvhvgounldxkdzndwybxhtycmlybhaaqvodntsvfhwcuhvuccwcsxelafyzushjhfyklvghpfvknprfouevsxmcuhiiiewcluehpmzrjzffnrptwbuhnyahrbzqvirvmffbxvrmynfcnupnukayjghpusewdwrbkhvjnveuiionefmnfxao", "eweweweweweweweweweweweweweueueueueueueueueueueueueueueuhuhuhuhuhuhshshshshshshshshshshshshshshshshshshshphphphpcpcpcpcpcpcpcpcpcpcpcpcpcpcpcrcrcrcrcrcrcrcrcrcrcrcrmrmrmrmrmrmrmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmgmgvgvgvgvgvgvgvgvgvgvgvgvgvgvgvgvovovovovovovovovonononononononononbnbnbnbnbnbnbnbnbnbnbnbnbnbnbabaiaiaiaiaiaiaiaiaiaiaiaiaiaiaiaiatatatatatftftftftftftftftftftftfdfdfdfdfdfdfdfdfdfdfdydydydydyzyzyzyzyzyzyzyzyzyzyzyzyzyzyjyjyjyjkjkjkjkjkjkjkjklklklklklklklklklklklkqkqkqwqwqwqwqwqwqwqwq"
  T(_t, S(i), S(Valid(reorganizeString(i), o)), S(o))
}

func TestReorganizeStringAppend(_t *testing.T) {
  i, o := "aab", "aba"
  T(_t, S(i), S(Valid(reorganizeStringAppend(i), o)), S(o))

  i, o = "aaab", ""
  T(_t, S(i), S(Valid(reorganizeStringAppend(i), o)), S(o))

  i, o = "zifrfbctby", "bibrftfycz"
  T(_t, S(i), S(Valid(reorganizeStringAppend(i), o)), S(o))

  i, o = "zqugrfbsznyiwbokwkpvpmeyvaosdkedbgjogzdpwawwl", "wzwzwzwawabebebsdsdvdvgygygfkikjklomonoqprpup"
  T(_t, S(i), S(Valid(reorganizeStringAppend(i), o)), S(o))

  i, o = "tndsewnllhrtwsvxenkscbivijfqnysamckzoyfnapuotmdexzkkrpmppttficzerdndssuveompqkemtbwbodrhwsfpbmkafpwyedpcowruntvymxtyyejqtajkcjakghtdwmuygecjncxzcxezgecrxonnszmqmecgvqqkdagvaaucewelchsmebikscciegzoiamovdojrmmwgbxeygibxxltemfgpogjkhobmhwquizuwvhfaiavsxhiknysdghcawcrphaykyashchyomklvghkyabxatmrkmrfsppfhgrwywtlxebgzmevefcqquvhvgounldxkdzndwybxhtycmlybhaaqvodntsvfhwcuhvuccwcsxelafyzushjhfyklvghpfvknprfouevsxmcuhiiiewcluehpmzrjzffnrptwbuhnyahrbzqvirvmffbxvrmynfcnupnukayjghpusewdwrbkhvjnveuiionefmnfxao", "eweweweweweweweweweweweweweueueueueueueueueueueueueueueuhuhuhuhuhuhshshshshshshshshshshshshshshshshshshshphphphpcpcpcpcpcpcpcpcpcpcpcpcpcpcpcrcrcrcrcrcrcrcrcrcrcrcrmrmrmrmrmrmrmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmxmgmgvgvgvgvgvgvgvgvgvgvgvgvgvgvgvgvovovovovovovovovonononononononononbnbnbnbnbnbnbnbnbnbnbnbnbnbnbabaiaiaiaiaiaiaiaiaiaiaiaiaiaiaiaiatatatatatftftftftftftftftftftftfdfdfdfdfdfdfdfdfdfdfdydydydydyzyzyzyzyzyzyzyzyzyzyzyzyzyzyjyjyjyjkjkjkjkjkjkjkjklklklklklklklklklklklkqkqkqwqwqwqwqwqwqwqwq"
  T(_t, S(i), S(Valid(reorganizeString(i), o)), S(o))
}

func Valid(s, o string) string {
  if len(s) == 0 {
    return ""
  }

  m := [26]int{}
  for _, v := range s {
    m[v - 'a']++
  }
  for _, v := range o {
    m[v - 'a']--
  }
  for _, v := range m {
    if v != 0 {
      return s
    }
  }

  for i := 0; i < len(s)-1; i++ {
    if s[i] == s[i+1] {
      return s
    }
  }

  return o
}
