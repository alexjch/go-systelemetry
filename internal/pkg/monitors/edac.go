package monitors

import (
	"fmt"

	"github.com/alexjch/go-dmesg/pkg/dmesg"
	"github.com/alexjch/go-systelemetry/internal/pkg/comm"
)

// EdacMonitor is a monitor for EDAC events and
// client for the systelem comm.Server when edac
// events are received these are forwarded to the
// comm.Server.
type EdacMonitor struct {
	comm.Client                   // embedded comm.Client
	*dmesg.Decoder                // embedded dmesg decoder
	scanner        *dmesg.Scanner // dmesg scanner
}

// follow is the callback function to forward EDAC
// events to the comm.Server.
func (e *EdacMonitor) follow(r *dmesg.Record) {
	// Notify the comm.Server
	res, err := e.Notify(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Response: %s\n", res.Msg)
}

// Listen starts the monitor and waits for EDAC events
// to be received.
func (e *EdacMonitor) Listen() error {
	if err := e.Follow(e.follow); err != nil {
		return err
	}
	return nil
}

// Close closes the scanner and stops the embedded
// dmesg decoder.
func (e *EdacMonitor) Close() error {
	err := e.scanner.Close()
	e.Stop()
	return err
}

// NewEdacMonitor initializes a new EdacMonitor and the
// underlaying dmesg scanner and the embedded decoder.
func NewEdacMonitor() (*EdacMonitor, error) {
	scanner, err := dmesg.NewScanner()
	if err != nil {
		return nil, err
	}
	decoder, err := dmesg.NewDecoder(scanner)
	if err != nil {
		return nil, err
	}
	return &EdacMonitor{
		scanner: scanner,
		Decoder: decoder,
	}, nil
}
