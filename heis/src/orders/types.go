package order


type ActionType int
type Direction int

const (
	Down Direction = -1
	Up             = 1
	Stop           = 0
)

const (
	UpButton 	int = iota
	DownButton
	PanelButton
)

type TenderType struct{
	time 	time.Time
	val 	int
}

const TakeActiveTender	 = 500 // Milliseconds
const TakeLostTender     = 20  // Seconds
const SamplingTime		 = 1   // Milliseconds
const Floors 			 = 4   // Number of floors
const Buttons			 = 3   // Number of buttons
const (
	InvalidMsg ActionType  = iota  //	Only used to check if the message recieved is of type ButtonMsg.
	NewOrder		 //
	DeleteOrder
	Tender
	AddOrder
)
type OrderType struct{
	Button 	int			// Holds the button on the floor, Up or Down
	Floor 	int			// Holds the floor
}

type OrderMsg struct {
	Action    	ActionType   	// Holds what the information of what to do with the message
	Order 		OrderType 		// Holds the floor and button of the order
	TenderVal 	int				// If the action is a Tender, this will hold the cost from the sender, that is, the value from the cost function for this order
}
