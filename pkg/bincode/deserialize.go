package bincode

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

func DeserializeData(data []byte, target interface{}) error {
	targetType := reflect.TypeOf(target)
	targetValue := reflect.ValueOf(target)

	// 如果 target 不是指针，返回错误
	if targetType.Kind() != reflect.Ptr {
		return fmt.Errorf("target must be a pointer")
	}

	// 创建一个目标类型的实例
	instance := reflect.New(targetType.Elem()).Elem()

	// 调用 deSerializeData 填充数据
	if err := deserializeData(data, instance); err != nil {
		return err
	}
	// 将填充好的实例赋值给 target
	targetValue.Elem().Set(instance)

	return nil
}

func MustDeserializeData(data []byte, target interface{}) {
	targetType := reflect.TypeOf(target)
	targetValue := reflect.ValueOf(target)

	// 如果 target 不是指针，返回错误
	if targetType.Kind() != reflect.Ptr {
		panic(fmt.Errorf("target must be a pointer"))
	}

	// 创建一个目标类型的实例
	instance := reflect.New(targetType.Elem()).Elem()

	// 调用 deSerializeData 填充数据
	if err := deserializeData(data, instance); err != nil {
		panic(err)
	}

	// 将填充好的实例赋值给 target
	targetValue.Elem().Set(instance)
}

func deserializeData(data []byte, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Bool:
		v.SetBool(data[0] != 0)
	case reflect.Uint8:
		v.SetUint(uint64(data[0]))
	case reflect.Int16:
		v.SetInt(int64(binary.LittleEndian.Uint16(data)))
	case reflect.Uint16:
		v.SetUint(uint64(binary.LittleEndian.Uint16(data)))
	case reflect.Int32:
		v.SetInt(int64(binary.LittleEndian.Uint32(data)))
	case reflect.Uint32:
		v.SetUint(uint64(binary.LittleEndian.Uint32(data)))
	case reflect.Int64:
		v.SetInt(int64(binary.LittleEndian.Uint64(data)))
	case reflect.Uint64:
		v.SetUint(binary.LittleEndian.Uint64(data))
	case reflect.Slice:
		switch v.Type().Elem().Kind() {
		case reflect.Array:
			length := int(binary.LittleEndian.Uint64(data[:8]))
			restData := data[8:]
			sliceType := reflect.SliceOf(v.Type().Elem())
			slice := reflect.MakeSlice(sliceType, length, length)
			elemType := v.Type().Elem()
			for i := 0; i < length; i++ {
				elemSize := elemType.Size()
				elemValue := reflect.New(elemType).Elem()
				if err := deserializeData(restData[:elemSize], elemValue); err != nil {
					return err
				}
				slice.Index(i).Set(elemValue)
				restData = restData[elemSize:]
			}
			v.Set(slice)
		}
	case reflect.Array:
		switch v.Type().Elem().Kind() {
		case reflect.Uint8:
			for i := 0; i < v.Len(); i++ {
				v.Index(i).SetUint(uint64(data[i]))
			}
		}
	case reflect.String:
		length := int(binary.LittleEndian.Uint64(data[:8]))
		strData := data[8 : 8+length]
		v.SetString(string(strData))
	case reflect.Ptr:
		if data[0] == 0 {
			v.Set(reflect.Zero(v.Type()))
		} else {
			elemValue := reflect.New(v.Type().Elem()).Elem()
			if err := deserializeData(data[1:], elemValue); err != nil {
				return err
			}
			v.Set(elemValue.Addr())
		}
	case reflect.Struct:
		var offset uintptr
		for i := 0; i < v.NumField(); i++ {
			fieldSize := v.Field(i).Type().Size()
			fieldData := data[offset : offset+fieldSize]
			fieldValue := reflect.New(v.Field(i).Type()).Elem()
			if err := deserializeData(fieldData, fieldValue); err != nil {
				return err
			}
			v.Field(i).Set(fieldValue)
			offset += fieldSize
		}
	}
	return nil
}
