// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"
)

// VisaApplicationContractABI is the input ABI used to generate the binding from.
const VisaApplicationContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"application\",\"outputs\":[{\"name\":\"destination\",\"type\":\"bytes32\"},{\"name\":\"arrivalDate\",\"type\":\"uint256\"},{\"name\":\"departureDate\",\"type\":\"uint256\"},{\"name\":\"submitted\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"submit\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"applicant\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"birthDate\",\"type\":\"uint256\"},{\"name\":\"applied\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"bytes32\"},{\"name\":\"arrivalDate\",\"type\":\"uint256\"},{\"name\":\"departureDate\",\"type\":\"uint256\"}],\"name\":\"apply\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"birthDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Updating\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Applying\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Applied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"}]"

// VisaApplicationContractBin is the compiled bytecode used for deploying new contracts.
const VisaApplicationContractBin = `0x608060405234801561001057600080fd5b506040516060806105e1833981016040908152815160208301519190920151821580159061003d57508115155b801561004857508015155b15156100db57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f506c6561736520737570706c7920616c6c20706572736f6e616c20646574616960448201527f6c732e0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008054600160a060020a031916331790556001929092556002556003556104d9806101086000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166326e4ca8281146100665780635bcb2fc6146100a357806388e687ea146100cc5780639eea2c7d14610125575b600080fd5b34801561007257600080fd5b5061007b610143565b6040805194855260208501939093528383019190915215156060830152519081900360800190f35b3480156100af57600080fd5b506100b8610155565b604080519115158252519081900360200190f35b3480156100d857600080fd5b506100e16102ef565b6040805173ffffffffffffffffffffffffffffffffffffffff9096168652602086019490945284840192909252606084015215156080830152519081900360a00190f35b34801561013157600080fd5b506100b860043560243560443561031d565b60055460065460075460085460ff1684565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461017757fe5b60045460ff16151561021057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4170706c7920666f722076697361206170706c69636174696f6e20666972737460448201527f2e00000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60085460ff16156102a857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f56697361206170706c69636174696f6e20616c7265616479207375626d69747460448201527f65642e0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6040805133815290517f1ed657d5540dc57b6181fb8e1b90fa1620897560baf17d1646f17d12ff256cb79181900360200190a1506008805460ff1916600190811790915590565b60005460015460025460035460045473ffffffffffffffffffffffffffffffffffffffff9094169360ff1685565b6000805473ffffffffffffffffffffffffffffffffffffffff16331461033f57fe5b60085460ff16156103d757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f56697361206170706c69636174696f6e20616c7265616479207375626d69747460448201527f65642e0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60045460ff1615156001141561041f576040805133815290517fca11c3ed479ab868b0a7ec6dc02af29de9105a9afd710836210c93a376108caa9181900360200190a1610453565b6040805133815290517f8c3edaffd9fe7aa204527d05e2e71e3f071e6758a6c5e9c1c652943cf2dcee499181900360200190a15b6005849055600683905560078290556040805133815290517f83f970044e0c4e2d8a4922431cb57df4cd6dc71f67b619eb184d105e3b6d917a9181900360200190a1506004805460ff1916600190811790915593925050505600a165627a7a72305820e4545add49c9c330229760fee873e12bbe950d494ec59723b726e38725ef14df0029`

// DeployVisaApplicationContract deploys a new Ethereum contract, binding an instance of VisaApplicationContract to it.
func DeployVisaApplicationContract(auth *bind.TransactOpts, backend bind.ContractBackend, name [32]byte, id *big.Int, birthDate *big.Int) (common.Address, *types.Transaction, *VisaApplicationContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VisaApplicationContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VisaApplicationContractBin), backend, name, id, birthDate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VisaApplicationContract{VisaApplicationContractCaller: VisaApplicationContractCaller{contract: contract}, VisaApplicationContractTransactor: VisaApplicationContractTransactor{contract: contract}, VisaApplicationContractFilterer: VisaApplicationContractFilterer{contract: contract}}, nil
}

// VisaApplicationContract is an auto generated Go binding around an Ethereum contract.
type VisaApplicationContract struct {
	VisaApplicationContractCaller     // Read-only binding to the contract
	VisaApplicationContractTransactor // Write-only binding to the contract
	VisaApplicationContractFilterer   // Log filterer for contract events
}

// VisaApplicationContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VisaApplicationContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VisaApplicationContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VisaApplicationContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VisaApplicationContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VisaApplicationContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VisaApplicationContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VisaApplicationContractSession struct {
	Contract     *VisaApplicationContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VisaApplicationContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VisaApplicationContractCallerSession struct {
	Contract *VisaApplicationContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// VisaApplicationContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VisaApplicationContractTransactorSession struct {
	Contract     *VisaApplicationContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// VisaApplicationContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VisaApplicationContractRaw struct {
	Contract *VisaApplicationContract // Generic contract binding to access the raw methods on
}

// VisaApplicationContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VisaApplicationContractCallerRaw struct {
	Contract *VisaApplicationContractCaller // Generic read-only contract binding to access the raw methods on
}

// VisaApplicationContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VisaApplicationContractTransactorRaw struct {
	Contract *VisaApplicationContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVisaApplicationContract creates a new instance of VisaApplicationContract, bound to a specific deployed contract.
func NewVisaApplicationContract(address common.Address, backend bind.ContractBackend) (*VisaApplicationContract, error) {
	contract, err := bindVisaApplicationContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VisaApplicationContract{VisaApplicationContractCaller: VisaApplicationContractCaller{contract: contract}, VisaApplicationContractTransactor: VisaApplicationContractTransactor{contract: contract}, VisaApplicationContractFilterer: VisaApplicationContractFilterer{contract: contract}}, nil
}

// NewVisaApplicationContractCaller creates a new read-only instance of VisaApplicationContract, bound to a specific deployed contract.
func NewVisaApplicationContractCaller(address common.Address, caller bind.ContractCaller) (*VisaApplicationContractCaller, error) {
	contract, err := bindVisaApplicationContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VisaApplicationContractCaller{contract: contract}, nil
}

// NewVisaApplicationContractTransactor creates a new write-only instance of VisaApplicationContract, bound to a specific deployed contract.
func NewVisaApplicationContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VisaApplicationContractTransactor, error) {
	contract, err := bindVisaApplicationContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VisaApplicationContractTransactor{contract: contract}, nil
}

// NewVisaApplicationContractFilterer creates a new log filterer instance of VisaApplicationContract, bound to a specific deployed contract.
func NewVisaApplicationContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VisaApplicationContractFilterer, error) {
	contract, err := bindVisaApplicationContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VisaApplicationContractFilterer{contract: contract}, nil
}

// bindVisaApplicationContract binds a generic wrapper to an already deployed contract.
func bindVisaApplicationContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VisaApplicationContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VisaApplicationContract *VisaApplicationContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VisaApplicationContract.Contract.VisaApplicationContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VisaApplicationContract *VisaApplicationContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VisaApplicationContract.Contract.VisaApplicationContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VisaApplicationContract *VisaApplicationContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VisaApplicationContract.Contract.VisaApplicationContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VisaApplicationContract *VisaApplicationContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VisaApplicationContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VisaApplicationContract *VisaApplicationContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VisaApplicationContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VisaApplicationContract *VisaApplicationContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VisaApplicationContract.Contract.contract.Transact(opts, method, params...)
}

// Applicant is a free data retrieval call binding the contract method 0x88e687ea.
//
// Solidity: function applicant() constant returns(addr address, name bytes32, id uint256, birthDate uint256, applied bool)
func (_VisaApplicationContract *VisaApplicationContractCaller) Applicant(opts *bind.CallOpts) (struct {
	Addr      common.Address
	Name      [32]byte
	Id        *big.Int
	BirthDate *big.Int
	Applied   bool
}, error) {
	ret := new(struct {
		Addr      common.Address
		Name      [32]byte
		Id        *big.Int
		BirthDate *big.Int
		Applied   bool
	})
	out := ret
	err := _VisaApplicationContract.contract.Call(opts, out, "applicant")
	return *ret, err
}

// Applicant is a free data retrieval call binding the contract method 0x88e687ea.
//
// Solidity: function applicant() constant returns(addr address, name bytes32, id uint256, birthDate uint256, applied bool)
func (_VisaApplicationContract *VisaApplicationContractSession) Applicant() (struct {
	Addr      common.Address
	Name      [32]byte
	Id        *big.Int
	BirthDate *big.Int
	Applied   bool
}, error) {
	return _VisaApplicationContract.Contract.Applicant(&_VisaApplicationContract.CallOpts)
}

// Applicant is a free data retrieval call binding the contract method 0x88e687ea.
//
// Solidity: function applicant() constant returns(addr address, name bytes32, id uint256, birthDate uint256, applied bool)
func (_VisaApplicationContract *VisaApplicationContractCallerSession) Applicant() (struct {
	Addr      common.Address
	Name      [32]byte
	Id        *big.Int
	BirthDate *big.Int
	Applied   bool
}, error) {
	return _VisaApplicationContract.Contract.Applicant(&_VisaApplicationContract.CallOpts)
}

// Application is a free data retrieval call binding the contract method 0x26e4ca82.
//
// Solidity: function application() constant returns(destination bytes32, arrivalDate uint256, departureDate uint256, submitted bool)
func (_VisaApplicationContract *VisaApplicationContractCaller) Application(opts *bind.CallOpts) (struct {
	Destination   [32]byte
	ArrivalDate   *big.Int
	DepartureDate *big.Int
	Submitted     bool
}, error) {
	ret := new(struct {
		Destination   [32]byte
		ArrivalDate   *big.Int
		DepartureDate *big.Int
		Submitted     bool
	})
	out := ret
	err := _VisaApplicationContract.contract.Call(opts, out, "application")
	return *ret, err
}

// Application is a free data retrieval call binding the contract method 0x26e4ca82.
//
// Solidity: function application() constant returns(destination bytes32, arrivalDate uint256, departureDate uint256, submitted bool)
func (_VisaApplicationContract *VisaApplicationContractSession) Application() (struct {
	Destination   [32]byte
	ArrivalDate   *big.Int
	DepartureDate *big.Int
	Submitted     bool
}, error) {
	return _VisaApplicationContract.Contract.Application(&_VisaApplicationContract.CallOpts)
}

// Application is a free data retrieval call binding the contract method 0x26e4ca82.
//
// Solidity: function application() constant returns(destination bytes32, arrivalDate uint256, departureDate uint256, submitted bool)
func (_VisaApplicationContract *VisaApplicationContractCallerSession) Application() (struct {
	Destination   [32]byte
	ArrivalDate   *big.Int
	DepartureDate *big.Int
	Submitted     bool
}, error) {
	return _VisaApplicationContract.Contract.Application(&_VisaApplicationContract.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0x9eea2c7d.
//
// Solidity: function apply(destination bytes32, arrivalDate uint256, departureDate uint256) returns(success bool)
func (_VisaApplicationContract *VisaApplicationContractTransactor) Apply(opts *bind.TransactOpts, destination [32]byte, arrivalDate *big.Int, departureDate *big.Int) (*types.Transaction, error) {
	return _VisaApplicationContract.contract.Transact(opts, "apply", destination, arrivalDate, departureDate)
}

// Apply is a paid mutator transaction binding the contract method 0x9eea2c7d.
//
// Solidity: function apply(destination bytes32, arrivalDate uint256, departureDate uint256) returns(success bool)
func (_VisaApplicationContract *VisaApplicationContractSession) Apply(destination [32]byte, arrivalDate *big.Int, departureDate *big.Int) (*types.Transaction, error) {
	return _VisaApplicationContract.Contract.Apply(&_VisaApplicationContract.TransactOpts, destination, arrivalDate, departureDate)
}

// Apply is a paid mutator transaction binding the contract method 0x9eea2c7d.
//
// Solidity: function apply(destination bytes32, arrivalDate uint256, departureDate uint256) returns(success bool)
func (_VisaApplicationContract *VisaApplicationContractTransactorSession) Apply(destination [32]byte, arrivalDate *big.Int, departureDate *big.Int) (*types.Transaction, error) {
	return _VisaApplicationContract.Contract.Apply(&_VisaApplicationContract.TransactOpts, destination, arrivalDate, departureDate)
}

// Submit is a paid mutator transaction binding the contract method 0x5bcb2fc6.
//
// Solidity: function submit() returns(success bool)
func (_VisaApplicationContract *VisaApplicationContractTransactor) Submit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VisaApplicationContract.contract.Transact(opts, "submit")
}

// Submit is a paid mutator transaction binding the contract method 0x5bcb2fc6.
//
// Solidity: function submit() returns(success bool)
func (_VisaApplicationContract *VisaApplicationContractSession) Submit() (*types.Transaction, error) {
	return _VisaApplicationContract.Contract.Submit(&_VisaApplicationContract.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0x5bcb2fc6.
//
// Solidity: function submit() returns(success bool)
func (_VisaApplicationContract *VisaApplicationContractTransactorSession) Submit() (*types.Transaction, error) {
	return _VisaApplicationContract.Contract.Submit(&_VisaApplicationContract.TransactOpts)
}

// VisaApplicationContractAppliedIterator is returned from FilterApplied and is used to iterate over the raw logs and unpacked data for Applied events raised by the VisaApplicationContract contract.
type VisaApplicationContractAppliedIterator struct {
	Event *VisaApplicationContractApplied // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VisaApplicationContractAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VisaApplicationContractApplied)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VisaApplicationContractApplied)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VisaApplicationContractAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VisaApplicationContractAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VisaApplicationContractApplied represents a Applied event raised by the VisaApplicationContract contract.
type VisaApplicationContractApplied struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApplied is a free log retrieval operation binding the contract event 0x83f970044e0c4e2d8a4922431cb57df4cd6dc71f67b619eb184d105e3b6d917a.
//
// Solidity: e Applied(sender address)
func (_VisaApplicationContract *VisaApplicationContractFilterer) FilterApplied(opts *bind.FilterOpts) (*VisaApplicationContractAppliedIterator, error) {

	logs, sub, err := _VisaApplicationContract.contract.FilterLogs(opts, "Applied")
	if err != nil {
		return nil, err
	}
	return &VisaApplicationContractAppliedIterator{contract: _VisaApplicationContract.contract, event: "Applied", logs: logs, sub: sub}, nil
}

// WatchApplied is a free log subscription operation binding the contract event 0x83f970044e0c4e2d8a4922431cb57df4cd6dc71f67b619eb184d105e3b6d917a.
//
// Solidity: e Applied(sender address)
func (_VisaApplicationContract *VisaApplicationContractFilterer) WatchApplied(opts *bind.WatchOpts, sink chan<- *VisaApplicationContractApplied) (event.Subscription, error) {

	logs, sub, err := _VisaApplicationContract.contract.WatchLogs(opts, "Applied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VisaApplicationContractApplied)
				if err := _VisaApplicationContract.contract.UnpackLog(event, "Applied", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// VisaApplicationContractApplyingIterator is returned from FilterApplying and is used to iterate over the raw logs and unpacked data for Applying events raised by the VisaApplicationContract contract.
type VisaApplicationContractApplyingIterator struct {
	Event *VisaApplicationContractApplying // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VisaApplicationContractApplyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VisaApplicationContractApplying)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VisaApplicationContractApplying)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VisaApplicationContractApplyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VisaApplicationContractApplyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VisaApplicationContractApplying represents a Applying event raised by the VisaApplicationContract contract.
type VisaApplicationContractApplying struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApplying is a free log retrieval operation binding the contract event 0x8c3edaffd9fe7aa204527d05e2e71e3f071e6758a6c5e9c1c652943cf2dcee49.
//
// Solidity: e Applying(sender address)
func (_VisaApplicationContract *VisaApplicationContractFilterer) FilterApplying(opts *bind.FilterOpts) (*VisaApplicationContractApplyingIterator, error) {

	logs, sub, err := _VisaApplicationContract.contract.FilterLogs(opts, "Applying")
	if err != nil {
		return nil, err
	}
	return &VisaApplicationContractApplyingIterator{contract: _VisaApplicationContract.contract, event: "Applying", logs: logs, sub: sub}, nil
}

// WatchApplying is a free log subscription operation binding the contract event 0x8c3edaffd9fe7aa204527d05e2e71e3f071e6758a6c5e9c1c652943cf2dcee49.
//
// Solidity: e Applying(sender address)
func (_VisaApplicationContract *VisaApplicationContractFilterer) WatchApplying(opts *bind.WatchOpts, sink chan<- *VisaApplicationContractApplying) (event.Subscription, error) {

	logs, sub, err := _VisaApplicationContract.contract.WatchLogs(opts, "Applying")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VisaApplicationContractApplying)
				if err := _VisaApplicationContract.contract.UnpackLog(event, "Applying", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// VisaApplicationContractSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the VisaApplicationContract contract.
type VisaApplicationContractSubmittedIterator struct {
	Event *VisaApplicationContractSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VisaApplicationContractSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VisaApplicationContractSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VisaApplicationContractSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VisaApplicationContractSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VisaApplicationContractSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VisaApplicationContractSubmitted represents a Submitted event raised by the VisaApplicationContract contract.
type VisaApplicationContractSubmitted struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x1ed657d5540dc57b6181fb8e1b90fa1620897560baf17d1646f17d12ff256cb7.
//
// Solidity: e Submitted(sender address)
func (_VisaApplicationContract *VisaApplicationContractFilterer) FilterSubmitted(opts *bind.FilterOpts) (*VisaApplicationContractSubmittedIterator, error) {

	logs, sub, err := _VisaApplicationContract.contract.FilterLogs(opts, "Submitted")
	if err != nil {
		return nil, err
	}
	return &VisaApplicationContractSubmittedIterator{contract: _VisaApplicationContract.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x1ed657d5540dc57b6181fb8e1b90fa1620897560baf17d1646f17d12ff256cb7.
//
// Solidity: e Submitted(sender address)
func (_VisaApplicationContract *VisaApplicationContractFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *VisaApplicationContractSubmitted) (event.Subscription, error) {

	logs, sub, err := _VisaApplicationContract.contract.WatchLogs(opts, "Submitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VisaApplicationContractSubmitted)
				if err := _VisaApplicationContract.contract.UnpackLog(event, "Submitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// VisaApplicationContractUpdatingIterator is returned from FilterUpdating and is used to iterate over the raw logs and unpacked data for Updating events raised by the VisaApplicationContract contract.
type VisaApplicationContractUpdatingIterator struct {
	Event *VisaApplicationContractUpdating // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VisaApplicationContractUpdatingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VisaApplicationContractUpdating)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VisaApplicationContractUpdating)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VisaApplicationContractUpdatingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VisaApplicationContractUpdatingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VisaApplicationContractUpdating represents a Updating event raised by the VisaApplicationContract contract.
type VisaApplicationContractUpdating struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdating is a free log retrieval operation binding the contract event 0xca11c3ed479ab868b0a7ec6dc02af29de9105a9afd710836210c93a376108caa.
//
// Solidity: e Updating(sender address)
func (_VisaApplicationContract *VisaApplicationContractFilterer) FilterUpdating(opts *bind.FilterOpts) (*VisaApplicationContractUpdatingIterator, error) {

	logs, sub, err := _VisaApplicationContract.contract.FilterLogs(opts, "Updating")
	if err != nil {
		return nil, err
	}
	return &VisaApplicationContractUpdatingIterator{contract: _VisaApplicationContract.contract, event: "Updating", logs: logs, sub: sub}, nil
}

// WatchUpdating is a free log subscription operation binding the contract event 0xca11c3ed479ab868b0a7ec6dc02af29de9105a9afd710836210c93a376108caa.
//
// Solidity: e Updating(sender address)
func (_VisaApplicationContract *VisaApplicationContractFilterer) WatchUpdating(opts *bind.WatchOpts, sink chan<- *VisaApplicationContractUpdating) (event.Subscription, error) {

	logs, sub, err := _VisaApplicationContract.contract.WatchLogs(opts, "Updating")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VisaApplicationContractUpdating)
				if err := _VisaApplicationContract.contract.UnpackLog(event, "Updating", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
