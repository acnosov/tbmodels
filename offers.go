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

var OfferMap = map[string]AB{
	//Correct Score
	"cs": {Kind: NoBet},
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
	//1 X 2 (reg. time)
	"time_win,tinnings,1,wdw": {Kind: NoBet},
	//1X2 + Both Score
	"mo_both_score": {Kind: NoBet},

	"wdw":          {Kind: WDW},
	"dc":           {Kind: DoubleChance},
	"ml":           {B: "for,ml,h", A: "for,ml,a", Kind: Moneyline},
	"oe":           {B: "for,odd", A: "for,even", Kind: Moneyline},
	"clean,a":      {B: "for,clean,a", A: "for,clean,a,no", Kind: Moneyline},
	"clean,h":      {B: "for,clean,h", A: "for,clean,h,no", Kind: Moneyline},
	"score,a":      {B: "for,score,a", A: "for,score,a,no", Kind: Moneyline},
	"score,h":      {B: "for,score,h", A: "for,score,h,no", Kind: Moneyline},
	"score,both":   {B: "for,score,both", A: "for,score,both,no", Kind: Moneyline},
	"win_to_nil,a": {B: "for,wintonil,a", A: "for,wintonil,a,no", Kind: Moneyline},
	"win_to_nil,h": {B: "for,wintonil,h", A: "for,wintonil,h,no", Kind: Moneyline},

	"time_win,tp,all,ml":  {B: "for,tp,all,ml,h", A: "for,tp,all,ml,a", Short: "ml", Kind: Moneyline},
	"time_win,tp,reg,ml":  {B: "for,tp,reg,ml,h", A: "for,tp,reg,ml,a", Short: "mlr", Kind: Moneyline},
	"time_win,thalf,1,ml": {B: "for,thalf,1,ml,h", A: "for,thalf,1,ml,a", Short: "ml1", Kind: Moneyline},

	"time_win,tperiod,1,ml": {B: "for,tp,1,ml,h", A: "for,tp,1,ml,a", Short: "ml1", Kind: Moneyline},
	"time_win,tperiod,2,ml": {B: "for,tp,2,ml,h", A: "for,tp,2,ml,a", Short: "ml2", Kind: Moneyline},
	"time_win,tperiod,3,ml": {B: "for,tp,3,ml,h", A: "for,tp,3,ml,a", Short: "ml3", Kind: Moneyline},
	"time_win,tperiod,4,ml": {B: "for,tp,4,ml,h", A: "for,tp,4,ml,a", Short: "ml4", Kind: Moneyline},

	"tennis_match,all": {B: "for,tset,all,vwhatever,p2", A: "for,tset,all,vwhatever,p1", Short: "ml", Kind: Moneyline},
	"tennis_match,1":   {B: "for,tset,1,vwhatever,p2", A: "for,tset,1,vwhatever,p1", Short: "ml1", Kind: Moneyline},
	"tennis_match,2":   {B: "for,tset,2,vwhatever,p2", A: "for,tset,2,vwhatever,p1", Short: "ml2", Kind: Moneyline},
	"tennis_match,3":   {B: "for,tset,3,vwhatever,p2", A: "for,tset,3,vwhatever,p1", Short: "ml3", Kind: Moneyline},
	"tennis_match,4":   {B: "for,tset,4,vwhatever,p2", A: "for,tset,4,vwhatever,p1", Short: "ml4", Kind: Moneyline},

	"ah":      {B: "ah,h", A: "ah,a", Short: "h", Kind: Handicap},
	"ahou":    {B: "ahunder", A: "ahover", Short: "t", Kind: Handicap},
	"tahou,h": {B: "tahunder,h", A: "tahover,h", Short: "th", Kind: Handicap},
	"tahou,a": {B: "tahunder,a", A: "tahover,a", Short: "ta", Kind: Handicap},

	"time_ah,tp,all":    {B: "tp,all,ah,h", A: "tp,all,ah,a", Short: "h", Kind: Handicap},
	"time_ah,tp,reg":    {B: "tp,reg,ah,h", A: "tp,reg,ah,a", Short: "hr", Kind: Handicap},
	"time_ah,tperiod,1": {B: "tp,1,ah,h", A: "tp,1,ah,a", Short: "h1", Kind: Handicap},
	"time_ah,tperiod,2": {B: "tp,2,ah,h", A: "tp,2,ah,a", Short: "h2", Kind: Handicap},
	"time_ah,tperiod,3": {B: "tp,3,ah,h", A: "tp,3,ah,h", Short: "h3", Kind: Handicap},
	"time_ah,tperiod,4": {B: "tp,4,ah,h", A: "tp,4,ah,h", Short: "h4", Kind: Handicap},

	"time_ah,thalf,1": {B: "thalf,1,ah,h", A: "thalf,1,ah,a", Short: "h1", Kind: Handicap},

	"time_ahou,tp,all":    {B: "tp,all,ahunder", A: "tp,all,ahover", Short: "t", Kind: Handicap},
	"time_ahou,tp,reg":    {B: "tp,reg,ahunder", A: "tp,reg,ahover", Short: "tr", Kind: Handicap},
	"time_ahou,tperiod,1": {B: "tp,1,ahunder", A: "tp,1,ahover", Short: "t1", Kind: Handicap},
	"time_ahou,tperiod,2": {B: "tp,2,ahunder", A: "tp,2,ahover", Short: "t2", Kind: Handicap},
	"time_ahou,tperiod,3": {B: "tp,3,ahunder", A: "tp,3,ahover", Short: "t3", Kind: Handicap},
	"time_ahou,tperiod,4": {B: "tp,4,ahunder", A: "tp,4,ahover", Short: "t4", Kind: Handicap},
	"time_ahou,thalf,1":   {B: "thalf,1,ahunder", A: "thalf,1,ahover", Short: "t1", Kind: Handicap},

	"tennis_ah,all,set":    {B: "tset,all,vwhatever,set,ah,p2", A: "tset,all,vwhatever,set,ah,p1", Short: "hs", Kind: Handicap},
	"tennis_ah,all,game":   {B: "tset,all,vwhatever,game,ah,p2", A: "tset,all,vwhatever,game,ah,p1", Short: "h", Kind: Handicap},
	"tennis_ah,1,game":     {B: "tset,1,vwhatever,game,ah,p2", A: "tset,1,vwhatever,game,ah,p1", Short: "h1s", Kind: Handicap},
	"tennis_ahou,all,set":  {B: "tset,all,vwhatever,set,ahunder", A: "tset,all,vwhatever,set,ahover", Short: "ts", Kind: Handicap},
	"tennis_ahou,all,game": {B: "tset,all,vwhatever,game,ahunder", A: "tset,all,vwhatever,game,ahover", Short: "t", Kind: Handicap},
	"tennis_ahou,1,game":   {B: "tset,1,vwhatever,game,ahunder", A: "tset,1,vwhatever,game,ahover", Short: "t1s", Kind: Handicap},

	//	New moneyline events
	"time_win,tmap,1,ml":     {Short: "mlm1", Kind: Moneyline},
	"time_win,tmap,2,ml":     {Short: "mlm2", Kind: Moneyline},
	"time_win,tmap,3,ml":     {Short: "mlm3", Kind: Moneyline},
	"time_win,tmap,4,ml":     {Short: "mlm4", Kind: Moneyline},
	"time_win,tmap,5,ml":     {Short: "mlm5", Kind: Moneyline},
	"time_win,tquarter,1,ml": {Short: "mlq1", Kind: Moneyline},

	//	New handicap events
	"time_ahou,tmap,1":        {Kind: Handicap},
	"time_ahou,tmap,2":        {Kind: Handicap},
	"time_ahou,tmap,3":        {Kind: Handicap},
	"time_ahou,tmap,4":        {Kind: Handicap},
	"time_ahou,tmap,5":        {Kind: Handicap},
	"time_ah,tmap,1":          {Kind: Handicap},
	"time_ah,tmap,2":          {Kind: Handicap},
	"time_ah,tmap,3":          {Kind: Handicap},
	"time_ah,tmap,4":          {Kind: Handicap},
	"time_ah,tmap,5":          {Kind: Handicap},
	"time_tahou,tmap,1,h":     {Kind: Handicap},
	"time_tahou,tmap,1,a":     {Kind: Handicap},
	"time_tahou,tmap,2,h":     {Kind: Handicap},
	"time_tahou,tmap,2,a":     {Kind: Handicap},
	"time_tahou,tmap,3,h":     {Kind: Handicap},
	"time_tahou,tmap,3,a":     {Kind: Handicap},
	"time_tahou,tmap,4,h":     {Kind: Handicap},
	"time_tahou,tmap,4,a":     {Kind: Handicap},
	"time_tahou,tmap,5,h":     {Kind: Handicap},
	"time_tahou,tmap,5,a":     {Kind: Handicap},
	"time_ah,tquarter,1":      {Kind: Handicap},
	"time_ahou,tquarter,1":    {Kind: Handicap},
	"time_tahou,thalf,1,a":    {Kind: Handicap},
	"time_tahou,thalf,1,h":    {Kind: Handicap},
	"time_tahou,tquarter,1,a": {Kind: Handicap},
	"time_tahou,tquarter,1,h": {Kind: Handicap},
	"time_tahou,tp,all,h":     {Kind: Handicap},
	"time_tahou,tp,all,a":     {Kind: Handicap},
	"time_tahou,tperiod,1,h":  {Kind: Handicap},
	"time_tahou,tperiod,2,h":  {Kind: Handicap},
	"time_tahou,tperiod,3,h":  {Kind: Handicap},
	"time_tahou,tperiod,1,a":  {Kind: Handicap},
	"time_tahou,tperiod,2,a":  {Kind: Handicap},
	"time_tahou,tperiod,3,a":  {Kind: Handicap},
	"time_tahou,tp,reg,a":     {Kind: Handicap},
	"time_tahou,tp,reg,h":     {Kind: Handicap},
}
