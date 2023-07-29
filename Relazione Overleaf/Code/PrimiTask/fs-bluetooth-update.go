func (fs _fsimpl) UpdateBluetoothMode(ctx context.Context, filename string, 
newMode string) (string, error) {
	var message string
	destinationFileName := path.Join(fs.cfg.Directory, "bluetoothconnections", 
  fmt.Sprintf("%s.json", filename))

	_, err := os.Stat(destinationFileName)
	if os.IsNotExist(err) {
		return types.ErrorFileNotFound, err
	}

	byteValue, err := os.ReadFile(destinationFileName)
	if err != nil {
		return message, fmt.Errorf("opening json file: %w", err)
	}

	var bluetoothStruct types.BluetoothConnectionsStruct
	err = json.Unmarshal(byteValue, &bluetoothStruct)
	if err != nil {
		return message, fmt.Errorf("unmarshal json file: %w", err)
	}
	*bluetoothStruct.Mode = newMode

	fp, err := os.Create(destinationFileName)
	if err != nil {
		return message, fmt.Errorf("truncating destination file: %w", err)
	}
	defer func() { _ = fp.Close() }()

	err = json.NewEncoder(fp).Encode(bluetoothStruct)
	if err != nil {
		return message, fmt.Errorf("converting connections to JSON: %w", err)
	}
	return message, nil
}