package govote

import (
    "errors"
) 

func init() {
    Schulze = SchulzeCtrl{}
    InstantRunoff = InstantRunoffCtrl{}
    Plurality = PluralityCtrl{}
}

type (
    // SchulzeCtrl is a struct exposed for reference -- use govote.Schulze instance
    SchulzeCtrl struct {}

    // InstantRunoffCtrl is a struct exposed for reference -- use govote.InstantRunoff instance
    InstantRunoffCtrl struct {}

    // PluralityCtrl is a struct exposed for reference -- use govote.Plurality instance
    PluralityCtrl struct {}
)

// New creates a new Schulze method poll
func (_ SchulzeCtrl) New(candidates []string) (SchulzePoll, error) {
    removeDuplicates(&candidates)
    if len(candidates) < 2 { return SchulzePoll{}, errors.New("not enough candidates") }
    return SchulzePoll{candidates: candidates}, nil
}

// New creates a new Instant Runoff poll
func (_ InstantRunoffCtrl) New(candidates []string) (InstantRunoffPoll, error) {
    removeDuplicates(&candidates)
    if len(candidates) < 2 { return InstantRunoffPoll{}, errors.New("not enough candidates") }
    return InstantRunoffPoll{candidates: candidates}, nil
}

// New creates a new Plurality poll
func (_ PluralityCtrl) New(candidates []string) (PluralityPoll, error) {
    removeDuplicates(&candidates)
    if len(candidates) < 2 { return PluralityPoll{}, errors.New("not enough candidates") }
    return PluralityPoll{candidates: candidates}, nil
}

var (
    // Schulze handler
    Schulze SchulzeCtrl 

    // InstantRunoff handler
    InstantRunoff InstantRunoffCtrl   

    // Plurality handler
    Plurality PluralityCtrl
)