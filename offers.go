package tbmodels

type AB struct {
	A, B, Short string
	Kind        uint8
}

const (
	Moneyline    = 1
	Handicap     = 2
	WDW          = 3
	DoubleChance = 4
	NoBet        = 5
)

// all - o (inc.Overtime)
// reg - r (regular time)
// tperiod -p (period)
// thalf -h (half)
// tmap -m (map)
// tquarter -q (quarter)
// tset -s (set)
var OfferMap = map[string]AB{
	//Correct Score
	"qualify": {Kind: NoBet},
	"cs":      {Kind: NoBet},
	//Total Between
	"gr": {Kind: NoBet},
	//Exact Total
	"exact_total": {Kind: NoBet},
	//Win by
	"wm": {Kind: NoBet},
	//1X2 and O/U
	"moou": {Kind: NoBet},
	//1X2+O
	"aou": {Kind: NoBet},
	//Set Score
	"tennis_cs,all,set": {Kind: NoBet},
	//1 X 2 (reg. time)
	"time_win,tp,reg,wdw": {Kind: NoBet},
	"time_win,tp,all,wdw": {Kind: NoBet},
	//1 X 2 (reg. time)
	"time_win,tinnings,1,wdw": {Kind: NoBet},
	"time_ah,tinnings,1":      {Kind: NoBet}, //todo: check if avail in web version
	"time_ahou,tinnings,1":    {Kind: NoBet}, //todo: check if avail in web version
	//1X2 + Both Score
	"mo_both_score":               {Kind: NoBet},
	"time_cs,tp,all":              {Kind: NoBet},
	"time_cs,thalf,1":             {Kind: NoBet},
	"time_win,tp,all,sub,180,wdw": {Kind: NoBet}, // darts wdw
	"time_win,tperiod,1,wdw":      {Kind: NoBet}, // ih wdw
	"time_win,thalf,1,wdw":        {Kind: NoBet}, // baseball wdw

	"wdw":          {Kind: WDW},
	"dc":           {Kind: DoubleChance},
	"oe":           {B: "for,odd", A: "for,even", Kind: Moneyline},
	"clean,a":      {B: "for,clean,a", A: "for,clean,a,no", Kind: Moneyline},
	"clean,h":      {B: "for,clean,h", A: "for,clean,h,no", Kind: Moneyline},
	"score,a":      {B: "for,score,a", A: "for,score,a,no", Kind: Moneyline},
	"score,h":      {B: "for,score,h", A: "for,score,h,no", Kind: Moneyline},
	"score,both":   {B: "for,score,both", A: "for,score,both,no", Kind: Moneyline},
	"win_to_nil,a": {B: "for,wintonil,a", A: "for,wintonil,a,no", Kind: Moneyline},
	"win_to_nil,h": {B: "for,wintonil,h", A: "for,wintonil,h,no", Kind: Moneyline},

	// moneyline and different periods (14)
	"ml":                     {B: "for,ml,h", A: "for,ml,a", Short: "ml", Kind: Moneyline},
	"time_win,tp,all,ml":     {B: "for,tp,all,ml,h", A: "for,tp,all,ml,a", Short: "mlo", Kind: Moneyline},
	"time_win,tp,reg,ml":     {B: "for,tp,reg,ml,h", A: "for,tp,reg,ml,a", Short: "mlr", Kind: Moneyline},
	"time_win,thalf,1,ml":    {B: "for,thalf,1,ml,h", A: "for,thalf,1,ml,a", Short: "ml1", Kind: Moneyline},
	"time_win,tquarter,1,ml": {B: "for,tquarter,1,ml,h", A: "for,tquarter,1,ml,a", Short: "mlq1", Kind: Moneyline},
	"time_win,tperiod,1,ml":  {B: "for,tp,1,ml,h", A: "for,tp,1,ml,a", Short: "mlp1", Kind: Moneyline},
	"time_win,tperiod,2,ml":  {B: "for,tp,2,ml,h", A: "for,tp,2,ml,a", Short: "mlp2", Kind: Moneyline},
	"time_win,tperiod,3,ml":  {B: "for,tp,3,ml,h", A: "for,tp,3,ml,a", Short: "mlp3", Kind: Moneyline},
	"time_win,tperiod,4,ml":  {B: "for,tp,4,ml,h", A: "for,tp,4,ml,a", Short: "mlp4", Kind: Moneyline},
	"time_win,tmap,1,ml":     {B: "for,tmap,1,ml,h", A: "for,tmap,1,ml,a", Short: "mlm1", Kind: Moneyline},
	"time_win,tmap,2,ml":     {B: "for,tmap,2,ml,h", A: "for,tmap,2,ml,a", Short: "mlm2", Kind: Moneyline},
	"time_win,tmap,3,ml":     {B: "for,tmap,3,ml,h", A: "for,tmap,3,ml,a", Short: "mlm3", Kind: Moneyline},
	"time_win,tmap,4,ml":     {B: "for,tmap,4,ml,h", A: "for,tmap,4,ml,a", Short: "mlm4", Kind: Moneyline},
	"time_win,tmap,5,ml":     {B: "for,tmap,5,ml,h", A: "for,tmap,5,ml,a", Short: "mlm5", Kind: Moneyline},

	// asian handicap and different periods (14)
	"ah":                 {B: "ah,h", A: "ah,a", Short: "h", Kind: Handicap},
	"time_ah,tp,all":     {B: "tp,all,ah,h", A: "tp,all,ah,a", Short: "ho", Kind: Handicap},
	"time_ah,tp,reg":     {B: "tp,reg,ah,h", A: "tp,reg,ah,a", Short: "hr", Kind: Handicap},
	"time_ah,thalf,1":    {B: "thalf,1,ah,h", A: "thalf,1,ah,a", Short: "h1", Kind: Handicap},
	"time_ah,tquarter,1": {B: "tquarter,1,ah,h", A: "tquarter,1,ah,a", Short: "hq1", Kind: Handicap},
	"time_ah,tperiod,1":  {B: "tp,1,ah,h", A: "tp,1,ah,a", Short: "hp1", Kind: Handicap},
	"time_ah,tperiod,2":  {B: "tp,2,ah,h", A: "tp,2,ah,a", Short: "hp2", Kind: Handicap},
	"time_ah,tperiod,3":  {B: "tp,3,ah,h", A: "tp,3,ah,h", Short: "hp3", Kind: Handicap},
	"time_ah,tperiod,4":  {B: "tp,4,ah,h", A: "tp,4,ah,h", Short: "hp4", Kind: Handicap},
	"time_ah,tmap,1":     {B: "tmap,1,ah,h", A: "tmap,1,ah,a", Short: "hm1", Kind: Handicap},
	"time_ah,tmap,2":     {B: "tmap,2,ah,h", A: "tmap,2,ah,a", Short: "hm2", Kind: Handicap},
	"time_ah,tmap,3":     {B: "tmap,3,ah,h", A: "tmap,3,ah,a", Short: "hm3", Kind: Handicap},
	"time_ah,tmap,4":     {B: "tmap,4,ah,h", A: "tmap,4,ah,a", Short: "hm4", Kind: Handicap},
	"time_ah,tmap,5":     {B: "tmap,5,ah,h", A: "tmap,5,ah,a", Short: "hm5", Kind: Handicap},
	// "time_ah,tinnings,1": {B: "tmap,5,ah,h", A: "tmap,5,ah,a", Short: "hm5", Kind: Handicap},
	// for,thalf,1,ah,h,-2
	//
	// asian total and different periods (14)
	"ahou":                     {B: "ahunder", A: "ahover", Short: "t", Kind: Handicap},
	"time_ahou,tp,all":         {B: "tp,all,ahunder", A: "tp,all,ahover", Short: "to", Kind: Handicap},
	"time_ahou,tp,all,sub,180": {B: "tp,all,sub,180,ahunder", A: "tp,all,sub,180,ahover", Short: "t180", Kind: Handicap},
	"time_ahou,tp,reg":         {B: "tp,reg,ahunder", A: "tp,reg,ahover", Short: "tr", Kind: Handicap},
	"time_ahou,thalf,1":        {B: "thalf,1,ahunder", A: "thalf,1,ahover", Short: "t1", Kind: Handicap},
	"time_ahou,tquarter,1":     {B: "tquarter,1,ahunder", A: "tquarter,1,ahover", Short: "tq1", Kind: Handicap},
	"time_ahou,tperiod,1":      {B: "tp,1,ahunder", A: "tp,1,ahover", Short: "tp1", Kind: Handicap},
	"time_ahou,tperiod,2":      {B: "tp,2,ahunder", A: "tp,2,ahover", Short: "tp2", Kind: Handicap},
	"time_ahou,tperiod,3":      {B: "tp,3,ahunder", A: "tp,3,ahover", Short: "tp3", Kind: Handicap},
	"time_ahou,tperiod,4":      {B: "tp,4,ahunder", A: "tp,4,ahover", Short: "tp4", Kind: Handicap},
	"time_ahou,tmap,1":         {B: "tmap,1,ahunder", A: "tmap,1,ahover", Short: "tm1", Kind: Handicap},
	"time_ahou,tmap,2":         {B: "tmap,2,ahunder", A: "tmap,2,ahover", Short: "tm2", Kind: Handicap},
	"time_ahou,tmap,3":         {B: "tmap,3,ahunder", A: "tmap,3,ahover", Short: "tm3", Kind: Handicap},
	"time_ahou,tmap,4":         {B: "tmap,4,ahunder", A: "tmap,4,ahover", Short: "tm4", Kind: Handicap},
	"time_ahou,tmap,5":         {B: "tmap,5,ahunder", A: "tmap,5,ahover", Short: "tm5", Kind: Handicap},
	// "time_ahou,tinnings,1":     {B: "tmap,5,ahunder", A: "tmap,5,ahover", Short: "tm5", Kind: Handicap},

	//	home total and different periods (14)
	"tahou,h":                 {B: "tahunder,h", A: "tahover,h", Short: "th", Kind: Handicap},
	"time_tahou,tp,all,h":     {B: "tp,all,tahunder,h", A: "tp,all,tahover,h", Short: "tho", Kind: Handicap},
	"time_tahou,tp,reg,h":     {B: "tp,reg,tahunder,h", A: "tp,reg,tahover,h", Short: "thr", Kind: Handicap},
	"time_tahou,thalf,1,h":    {B: "thalf,1,tahunder,h", A: "thalf,1,tahover,h", Short: "th1", Kind: Handicap},
	"time_tahou,tquarter,1,h": {B: "tquarter,1,tahunder,h", A: "tquarter,1,tahover,h", Short: "thq1", Kind: Handicap},
	"time_tahou,tperiod,1,h":  {B: "tp,1,tahunder,h", A: "tp,1,tahover,h", Short: "thp1", Kind: Handicap},
	"time_tahou,tperiod,2,h":  {B: "tp,2,tahunder,h", A: "tp,2,tahover,h", Short: "thp2", Kind: Handicap},
	"time_tahou,tperiod,3,h":  {B: "tp,3,tahunder,h", A: "tp,3,tahover,h", Short: "thp3", Kind: Handicap},
	"time_tahou,tperiod,4,h":  {B: "tp,4,tahunder,h", A: "tp,4,tahover,h", Short: "thp4", Kind: Handicap},
	"time_tahou,tmap,1,h":     {B: "tmap,1,tahunder,h", A: "tmap,1,tahover,h", Short: "thm1", Kind: Handicap},
	"time_tahou,tmap,2,h":     {B: "tmap,2,tahunder,h", A: "tmap,2,tahover,h", Short: "thm2", Kind: Handicap},
	"time_tahou,tmap,3,h":     {B: "tmap,3,tahunder,h", A: "tmap,3,tahover,h", Short: "thm3", Kind: Handicap},
	"time_tahou,tmap,4,h":     {B: "tmap,4,tahunder,h", A: "tmap,4,tahover,h", Short: "thm4", Kind: Handicap},
	"time_tahou,tmap,5,h":     {B: "tmap,5,tahunder,h", A: "tmap,5,tahover,h", Short: "thm5", Kind: Handicap},

	// away total and different periods (14)
	"tahou,a":                 {B: "tahunder,a", A: "tahover,a", Short: "ta", Kind: Handicap},
	"time_tahou,tp,all,a":     {B: "tp,all,tahunder,a", A: "tp,all,tahover,a", Short: "tao", Kind: Handicap},
	"time_tahou,tp,reg,a":     {B: "tp,reg,tahunder,a", A: "tp,reg,tahover,a", Short: "tar", Kind: Handicap},
	"time_tahou,thalf,1,a":    {B: "thalf,1,tahunder,a", A: "thalf,1,tahover,a", Short: "ta1", Kind: Handicap},
	"time_tahou,tquarter,1,a": {B: "tquarter,1,tahunder,a", A: "tquarter,1,tahover,a", Short: "taq1", Kind: Handicap},
	"time_tahou,tperiod,1,a":  {B: "tp,1,tahunder,a", A: "tp,1,tahover,a", Short: "tap1", Kind: Handicap},
	"time_tahou,tperiod,2,a":  {B: "tp,2,tahunder,a", A: "tp,2,tahover,a", Short: "tap2", Kind: Handicap},
	"time_tahou,tperiod,3,a":  {B: "tp,3,tahunder,a", A: "tp,3,tahover,a", Short: "tap3", Kind: Handicap},
	"time_tahou,tperiod,4,a":  {B: "tp,4,tahunder,a", A: "tp,4,tahover,a", Short: "tap4", Kind: Handicap},
	"time_tahou,tmap,1,a":     {B: "tmap,1,tahunder,a", A: "tmap,1,tahover,a", Short: "tam1", Kind: Handicap},
	"time_tahou,tmap,2,a":     {B: "tmap,2,tahunder,a", A: "tmap,2,tahover,a", Short: "tam2", Kind: Handicap},
	"time_tahou,tmap,3,a":     {B: "tmap,3,tahunder,a", A: "tmap,3,tahover,a", Short: "tam3", Kind: Handicap},
	"time_tahou,tmap,4,a":     {B: "tmap,4,tahunder,a", A: "tmap,4,tahover,a", Short: "tam4", Kind: Handicap},
	"time_tahou,tmap,5,a":     {B: "tmap,5,tahunder,a", A: "tmap,5,tahover,a", Short: "tam5", Kind: Handicap},

	// tennis match and periods
	// this market allows you to bet on the outright winner of the entire match, regardless of the number of sets played.
	"tennis_match,all": {B: "for,tset,all,vwhatever,p2", A: "for,tset,all,vwhatever,p1", Short: "ml", Kind: Moneyline},
	// bet on the winner of the set 1 in the match. (Set 1 Winner)
	"tennis_match,1": {B: "for,tset,1,vwhatever,p2", A: "for,tset,1,vwhatever,p1", Short: "ml1", Kind: Moneyline},
	// bet on the winner of the set 2 in the match. (Set 2 Winner)
	"tennis_match,2": {B: "for,tset,2,vwhatever,p2", A: "for,tset,2,vwhatever,p1", Short: "ml2", Kind: Moneyline},
	// bet on the winner of the set 3 in the match. (Set 3 Winner)
	"tennis_match,3": {B: "for,tset,3,vwhatever,p2", A: "for,tset,3,vwhatever,p1", Short: "ml3", Kind: Moneyline},
	// bet on the winner of the set 4 in the match. (Set 4 Winner)
	"tennis_match,4": {B: "for,tset,4,vwhatever,p2", A: "for,tset,4,vwhatever,p1", Short: "ml4", Kind: Moneyline},
	// bet on the winner of the set 5 in the match. (Set 5 Winner)
	"tennis_match,5": {B: "for,tset,5,vwhatever,p2", A: "for,tset,5,vwhatever,p1", Short: "ml5", Kind: Moneyline},
	// tennis handicap and totals
	// in this market, handicaps are applied to the total number of games won throughout the entire match.
	"tennis_ah,all,game": {B: "tset,all,vwhatever,game,ah,p2", A: "tset,all,vwhatever,game,ah,p1", Short: "h", Kind: Handicap},
	// Handicaps are applied to the total number of sets won by each player or team across the entire match.
	"tennis_ah,all,set": {B: "tset,all,vwhatever,set,ah,p2", A: "tset,all,vwhatever,set,ah,p1", Short: "hs", Kind: Handicap},

	// handicaps are applied to the total number of games won in the 1st set.
	"tennis_ah,1,game": {B: "tset,1,vwhatever,game,ah,p2", A: "tset,1,vwhatever,game,ah,p1", Short: "h1s", Kind: Handicap},
	"tennis_ah,2,game": {B: "tset,2,vwhatever,game,ah,p2", A: "tset,2,vwhatever,game,ah,p1", Short: "h2s", Kind: Handicap},
	"tennis_ah,3,game": {B: "tset,3,vwhatever,game,ah,p2", A: "tset,3,vwhatever,game,ah,p1", Short: "h3s", Kind: Handicap},
	"tennis_ah,4,game": {B: "tset,4,vwhatever,game,ah,p2", A: "tset,4,vwhatever,game,ah,p1", Short: "h4s", Kind: Handicap},
	"tennis_ah,5,game": {B: "tset,5,vwhatever,game,ah,p2", A: "tset,5,vwhatever,game,ah,p1", Short: "h5s", Kind: Handicap},
	"tennis_ah,6,game": {B: "tset,6,vwhatever,game,ah,p2", A: "tset,6,vwhatever,game,ah,p1", Short: "h6s", Kind: Handicap},
	"tennis_ah,7,game": {B: "tset,7,vwhatever,game,ah,p2", A: "tset,7,vwhatever,game,ah,p1", Short: "h7s", Kind: Handicap},

	// This market involves predicting whether the total number of games played in the first set will be over or under a specified number.
	"tennis_ahou,1,game": {B: "tset,1,vwhatever,game,ahunder", A: "tset,1,vwhatever,game,ahover", Short: "t1s", Kind: Handicap},
	"tennis_ahou,2,game": {B: "tset,2,vwhatever,game,ahunder", A: "tset,2,vwhatever,game,ahover", Short: "t2s", Kind: Handicap},
	"tennis_ahou,3,game": {B: "tset,3,vwhatever,game,ahunder", A: "tset,3,vwhatever,game,ahover", Short: "t3s", Kind: Handicap},
	"tennis_ahou,4,game": {B: "tset,4,vwhatever,game,ahunder", A: "tset,4,vwhatever,game,ahover", Short: "t4s", Kind: Handicap},
	"tennis_ahou,5,game": {B: "tset,5,vwhatever,game,ahunder", A: "tset,5,vwhatever,game,ahover", Short: "t5s", Kind: Handicap},
	"tennis_ahou,6,game": {B: "tset,6,vwhatever,game,ahunder", A: "tset,6,vwhatever,game,ahover", Short: "t6s", Kind: Handicap},
	"tennis_ahou,7,game": {B: "tset,7,vwhatever,game,ahunder", A: "tset,7,vwhatever,game,ahover", Short: "t7s", Kind: Handicap},
	// asian over/under (all sets)
	"tennis_ahou,all,set": {B: "tset,all,vwhatever,set,ahunder", A: "tset,all,vwhatever,set,ahover", Short: "ts", Kind: Handicap},
	// bettors predict the total number of games played throughout the entire match, with Asian handicap options available.
	"tennis_ahou,all,game": {B: "tset,all,vwhatever,game,ahunder", A: "tset,all,vwhatever,game,ahover", Short: "t", Kind: Handicap},
}
