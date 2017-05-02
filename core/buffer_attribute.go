/*
 * Ported from three.js by @rydrman
 */

package core

import "github.com/rydrman/three.go/math3"

type BufferAttribute struct {
	UUID string

	Array      []interface{}
	ItemSize   int
	Count      int
	Normalized bool

	Dynamic           bool
	UpdateRangeOffset int
	UpdateRangeCount  int

	onUploadCallback func()

	Version int
}

func NewBufferGeometry(array []interface{}, itemSize int, normalized bool) *BufferAttribute {

	return &BufferAttribute{

		UUID: math3.GenerateUUID(),

		Array:      array,
		ItemSize:   itemSize,
		Count:      len(array),
		Normalized: normalized,

		Dynamic:           false,
		UpdateRangeOffset: 0,
		UpdateRangeCount:  -1,

		onUploadCallback: nil,

		Version: 0,
	}

}

/*Object.defineProperty( BufferAttribute.prototype, 'needsUpdate', {

	func (attr *BufferAttribute) Set( value ) {

		if ( value === true ) attr.Version ++;

	}

} );

Object.assign( BufferAttribute.prototype, {

	isBufferAttribute: true,

	func (attr *BufferAttribute) SetArray( array ) {

		if ( Array.isArray( array ) ) {

			throw new TypeError( 'THREE.BufferAttribute: array should be a Typed Array.' );

		}

		attr.Count = array !== undefined ? array.length / attr.ItemSize : 0;
		attr.Array = array;

	},

	func (attr *BufferAttribute) SetDynamic( value ) {

		attr.Dynamic = value;

		return this;

	},

	func (attr *BufferAttribute) Copy( source ) {

		attr.Array = new source.array.constructor( source.array );
		attr.ItemSize = source.itemSize;
		attr.Count = source.count;
		attr.Normalized = source.normalized;

		attr.Dynamic = source.dynamic;

		return this;

	},

	func (attr *BufferAttribute) CopyAt( index1, attribute, index2 ) {

		index1 *= attr.ItemSize;
		index2 *= attribute.itemSize;

		for ( var i = 0, l = attr.ItemSize; i < l; i ++ ) {

			attr.Array[ index1 + i ] = attribute.array[ index2 + i ];

		}

		return this;

	},

	func (attr *BufferAttribute) CopyArray( array ) {

		attr.Array.set( array );

		return this;

	},

	func (attr *BufferAttribute) CopyColorsArray( colors ) {

		var array = attr.Array, offset = 0;

		for ( var i = 0, l = colors.length; i < l; i ++ ) {

			var color = colors[ i ];

			if ( color === undefined ) {

				console.warn( 'THREE.BufferAttribute.copyColorsArray(): color is undefined', i );
				color = new Color();

			}

			array[ offset ++ ] = color.r;
			array[ offset ++ ] = color.g;
			array[ offset ++ ] = color.b;

		}

		return this;

	},

	func (attr *BufferAttribute) CopyIndicesArray( indices ) {

		var array = attr.Array, offset = 0;

		for ( var i = 0, l = indices.length; i < l; i ++ ) {

			var index = indices[ i ];

			array[ offset ++ ] = index.a;
			array[ offset ++ ] = index.b;
			array[ offset ++ ] = index.c;

		}

		return this;

	},

	func (attr *BufferAttribute) CopyVector2sArray( vectors ) {

		var array = attr.Array, offset = 0;

		for ( var i = 0, l = vectors.length; i < l; i ++ ) {

			var vector = vectors[ i ];

			if ( vector === undefined ) {

				console.warn( 'THREE.BufferAttribute.copyVector2sArray(): vector is undefined', i );
				vector = new Vector2();

			}

			array[ offset ++ ] = vector.x;
			array[ offset ++ ] = vector.y;

		}

		return this;

	},

	func (attr *BufferAttribute) CopyVector3sArray( vectors ) {

		var array = attr.Array, offset = 0;

		for ( var i = 0, l = vectors.length; i < l; i ++ ) {

			var vector = vectors[ i ];

			if ( vector === undefined ) {

				console.warn( 'THREE.BufferAttribute.copyVector3sArray(): vector is undefined', i );
				vector = new Vector3();

			}

			array[ offset ++ ] = vector.x;
			array[ offset ++ ] = vector.y;
			array[ offset ++ ] = vector.z;

		}

		return this;

	},

	func (attr *BufferAttribute) CopyVector4sArray( vectors ) {

		var array = attr.Array, offset = 0;

		for ( var i = 0, l = vectors.length; i < l; i ++ ) {

			var vector = vectors[ i ];

			if ( vector === undefined ) {

				console.warn( 'THREE.BufferAttribute.copyVector4sArray(): vector is undefined', i );
				vector = new Vector4();

			}

			array[ offset ++ ] = vector.x;
			array[ offset ++ ] = vector.y;
			array[ offset ++ ] = vector.z;
			array[ offset ++ ] = vector.w;

		}

		return this;

	},

	func (attr *BufferAttribute) Set( value, offset ) {

		if ( offset === undefined ) offset = 0;

		attr.Array.set( value, offset );

		return this;

	},

	func (attr *BufferAttribute) GetX( index ) {

		return attr.Array[ index * attr.ItemSize ];

	},

	func (attr *BufferAttribute) SetX( index, x ) {

		attr.Array[ index * attr.ItemSize ] = x;

		return this;

	},

	func (attr *BufferAttribute) GetY( index ) {

		return attr.Array[ index * attr.ItemSize + 1 ];

	},

	func (attr *BufferAttribute) SetY( index, y ) {

		attr.Array[ index * attr.ItemSize + 1 ] = y;

		return this;

	},

	func (attr *BufferAttribute) GetZ( index ) {

		return attr.Array[ index * attr.ItemSize + 2 ];

	},

	func (attr *BufferAttribute) SetZ( index, z ) {

		attr.Array[ index * attr.ItemSize + 2 ] = z;

		return this;

	},

	func (attr *BufferAttribute) GetW( index ) {

		return attr.Array[ index * attr.ItemSize + 3 ];

	},

	func (attr *BufferAttribute) SetW( index, w ) {

		attr.Array[ index * attr.ItemSize + 3 ] = w;

		return this;

	},

	func (attr *BufferAttribute) SetXY( index, x, y ) {

		index *= attr.ItemSize;

		attr.Array[ index + 0 ] = x;
		attr.Array[ index + 1 ] = y;

		return this;

	},

	func (attr *BufferAttribute) SetXYZ( index, x, y, z ) {

		index *= attr.ItemSize;

		attr.Array[ index + 0 ] = x;
		attr.Array[ index + 1 ] = y;
		attr.Array[ index + 2 ] = z;

		return this;

	},

	func (attr *BufferAttribute) SetXYZW( index, x, y, z, w ) {

		index *= attr.ItemSize;

		attr.Array[ index + 0 ] = x;
		attr.Array[ index + 1 ] = y;
		attr.Array[ index + 2 ] = z;
		attr.Array[ index + 3 ] = w;

		return this;

	},

	func (attr *BufferAttribute) OnUpload( callback ) {

		attr.OnUploadCallback = callback;

		return this;

	},

	func (attr *BufferAttribute) Clone() {

		return new attr.Constructor( attr.Array, attr.ItemSize ).copy( this );

	}

} );

//

function Int8BufferAttribute( array, itemSize ) {

	BufferAttribute.call( this, new Int8Array( array ), itemSize );

}

Int8BufferAttribute.prototype = Object.create( BufferAttribute.prototype );
Int8BufferAttribute.prototype.constructor = Int8BufferAttribute;


function Uint8BufferAttribute( array, itemSize ) {

	BufferAttribute.call( this, new Uint8Array( array ), itemSize );

}

Uint8BufferAttribute.prototype = Object.create( BufferAttribute.prototype );
Uint8BufferAttribute.prototype.constructor = Uint8BufferAttribute;


function Uint8ClampedBufferAttribute( array, itemSize ) {

	BufferAttribute.call( this, new Uint8ClampedArray( array ), itemSize );

}

Uint8ClampedBufferAttribute.prototype = Object.create( BufferAttribute.prototype );
Uint8ClampedBufferAttribute.prototype.constructor = Uint8ClampedBufferAttribute;


function Int16BufferAttribute( array, itemSize ) {

	BufferAttribute.call( this, new Int16Array( array ), itemSize );

}

Int16BufferAttribute.prototype = Object.create( BufferAttribute.prototype );
Int16BufferAttribute.prototype.constructor = Int16BufferAttribute;


function Uint16BufferAttribute( array, itemSize ) {

	BufferAttribute.call( this, new Uint16Array( array ), itemSize );

}

Uint16BufferAttribute.prototype = Object.create( BufferAttribute.prototype );
Uint16BufferAttribute.prototype.constructor = Uint16BufferAttribute;


function Int32BufferAttribute( array, itemSize ) {

	BufferAttribute.call( this, new Int32Array( array ), itemSize );

}

Int32BufferAttribute.prototype = Object.create( BufferAttribute.prototype );
Int32BufferAttribute.prototype.constructor = Int32BufferAttribute;


function Uint32BufferAttribute( array, itemSize ) {

	BufferAttribute.call( this, new Uint32Array( array ), itemSize );

}

Uint32BufferAttribute.prototype = Object.create( BufferAttribute.prototype );
Uint32BufferAttribute.prototype.constructor = Uint32BufferAttribute;


function Float32BufferAttribute( array, itemSize ) {

	BufferAttribute.call( this, new Float32Array( array ), itemSize );

}

Float32BufferAttribute.prototype = Object.create( BufferAttribute.prototype );
Float32BufferAttribute.prototype.constructor = Float32BufferAttribute;


function Float64BufferAttribute( array, itemSize ) {

	BufferAttribute.call( this, new Float64Array( array ), itemSize );

}

Float64BufferAttribute.prototype = Object.create( BufferAttribute.prototype );
Float64BufferAttribute.prototype.constructor = Float64BufferAttribute;

//

export {
	Float64BufferAttribute,
	Float32BufferAttribute,
	Uint32BufferAttribute,
	Int32BufferAttribute,
	Uint16BufferAttribute,
	Int16BufferAttribute,
	Uint8ClampedBufferAttribute,
	Uint8BufferAttribute,
	Int8BufferAttribute,
	BufferAttribute
};
*/
